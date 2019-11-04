package main

//多个读写
import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"
)
//假设我们有一个读取器和一个写入器，读取器必须依赖写入器对缓冲区进行数据写入后，才可以从缓冲区中读取数据，
// 写入器每次完成写入数据后，都需要通过某种通知机制通知处于阻塞状态的读取器，告诉它可以对数据进行访问

// 数据 bucket
type DataBucket struct {
	buffer *bytes.Buffer  //缓冲区
	mutex *sync.RWMutex //互斥锁
	cond  *sync.Cond //条件变量
}

func NewDataBucket() *DataBucket {
	buf := make([]byte, 0)
	db := &DataBucket{
		buffer:     bytes.NewBuffer(buf),
		mutex: new(sync.RWMutex),
	}
	db.cond = sync.NewCond(db.mutex.RLocker())
	return db
}

// 读取器
func (db *DataBucket) Read(i int) {
	db.mutex.RLock()   // 打开读锁
	defer db.mutex.RUnlock()  // 结束后释放读锁
	var data []byte
	var d byte
	var err error
	for {
		//每次读取一个字节
		if d, err = db.buffer.ReadByte(); err != nil {
			if err == io.EOF { // 缓冲区数据为空时执行
				if string(data) != "" {  // data 不为空，则打印它
					fmt.Printf("reader-%d: %s\n", i, data)
				}
				db.cond.Wait() // 缓冲区为空，通过 Wait 方法等待通知，进入阻塞状态
				data = data[:0]  // 将 data 清空
				continue
			}
		}
		data = append(data, d) // 将读取到的数据添加到 data 中
	}
}

// 写入器
func (db *DataBucket) Put(d []byte) (int, error) {
	db.mutex.Lock()   // 打开写锁
	defer db.mutex.Unlock()  // 结束后释放写锁
	//写入一个数据块
	n, err := db.buffer.Write(d)
	db.cond.Broadcast()  // 写入数据后通过 Broadcast 通知处于阻塞状态的读取器
	return n, err
}

//使用了读写互斥锁，在读取器里面使用读锁，在写入器里面使用写锁，并且通过 defer 语句释放锁，然后在锁保护的情况下，
// 通过条件变量协调读写线程：在读线程中，当缓冲区为空的时候，通过 db.cond.Wait() 阻塞读线程；
// 在写线程中，当缓冲区写入数据的时候通过 db.cond.Signal() 通知读线程继续读取数据

//通知单个阻塞线程用  Signal 方法，通知多个阻塞线程需要使用 Broadcast 方法

func main() {
	db := NewDataBucket()
	for i := 1; i < 3; i++ {  // 启动多个读取器
		go db.Read(i)
	}
	for j := 0; j < 10; j++  {  // 启动多个写入器
		go func(i int) {
			d := fmt.Sprintf("data-%d", i)
			db.Put([]byte(d))  // 写入数据到缓冲区
		}(j)
		time.Sleep(100 * time.Millisecond) // 每次启动一个写入器暂停100ms，让读取器阻塞
	}
}
