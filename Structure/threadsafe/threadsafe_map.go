package main

import (
	"fmt"
	"sync"
	"time"
)

//map映射，瞬间查找
//线程安全的map结构，
//一边读一边写的时候

type SyncMap struct {
	mymap map[string]string  //字符串映射字符串
	*sync.RWMutex  //读写锁
}

var smap SyncMap
var done chan bool  //代表是否完成的通道

func write1()  {
	keys := []string{"1","2","3"}
	for _,k:= range keys{
		smap.Lock()
		smap.mymap[k] = k  //赋值
		smap.Unlock()
		time.Sleep(1*time.Second)
	}
	done <- true  //通道写入，
}

func write2()  {
	keys := []string{"a1","b2","c3"}
	for _,k:= range keys{
		smap.Lock()
		smap.mymap[k] = k  //赋值
		smap.Unlock()
		time.Sleep(1*time.Second)
	}
	done <- true  //通道写入，
}

func read1()  {
	smap.RLock()  //读的时候加锁
	fmt.Println("readlock")
	for k,v := range smap.mymap{
		fmt.Println(k,v)
	}
	smap.RUnlock()

}

func main()  {
	smap = SyncMap{make(map[string]string),new(sync.RWMutex)} //初始化
	done = make(chan bool,2)  //因为两个写入，所以给2个管道
	go write1()
	go write2()

	for {
		read1()
		if len(done) == 2 {
			fmt.Println(smap.mymap)
			break
		}else {
			time.Sleep(1*time.Second)
		}
	}
}
/*
运行结果：
readlock
readlock
1 1
a1 a1
readlock
1 1
a1 a1
2 2
b2 b2
readlock
a1 a1
2 2
b2 b2
c3 c3
3 3
1 1
map[1:1 2:2 3:3 a1:a1 b2:b2 c3:c3]
 */