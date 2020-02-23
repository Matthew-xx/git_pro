package pipelineMiddleware
//中间件

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)


var starttime time.Time  //构造时间
//初始化
func Init()  {
	starttime = time.Now()

}

func UserTime()  {
	fmt.Println(time.Since(starttime))  //统计消耗时间
}


//内存排序
func InMemorySort(in <- chan int) <- chan int {
	//将一个int管道写入in并返回一个管道
	out := make(chan int,1024)  //新管道
	go func() {
		data := []int{}  //创建数组存储数据并排序
		for v:= range in{
			data = append(data,v) //将数据压入数组
		}
		fmt.Println("数据读取完成,耗时",time.Since(starttime))
		//排序，可替换自定义排序函数
		sort.Ints(data)

		for _,v:= range data{
			out <- v  //压入数据
		}
		close(out)  //关闭管道
	}()  //构造线程，接收数据并排好序后再传入新管道，此时便是有序

	return out
}
//合并，两个管道的数据有序，归并为有序的数据并压入到另一个管道
func Merge(in1,in2 <- chan int) <-chan int {
	out := make(chan int,1024)
	go func() {
		v1,ok1 := <- in1
		v2,ok2 := <- in2

		//归并排序
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2){
				out <- v1  //v1小于v2，取出v1压入再次读取v1
				v1,ok1 = <- in1
			}else {
				out <- v2
				v2,ok2 = <- in2
			}
		}
		close(out)
	}()
	return out
}

//写入(将管道写入文件
func WriteSlink(writer io.Writer,in <- chan int)  {
	for v:= range in{
		buf := make([]byte,8)  //64位，8字节
		binary.BigEndian.PutUint64(buf,uint64(v))  //字节转换
		writer.Write(buf)  //写入
	}
}

//随机数数组
func RandSource(count int) <- chan int {
	out := make(chan int)
	go func() {
		for i:=0;i<count;i++ {
			out <- rand.Int()  //压入随机数
		}
		close(out)
	}()
	return out
}

//多路合并(多个输入，一个输出
func MergeN(inputs... <- chan int) <- chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}else {
		m := len(inputs)/2
		//一直递归
		return Merge(MergeN(inputs[:m]...),MergeN(inputs[:m]...))
	}
}

//读取数据
func ReaderSource(reader io.Reader,chunksize int) <- chan int {
	out := make(chan int,1024)
	go func() {
		buf := make([]byte,8)
		readersize := 0
		for {
			n,err := reader.Read(buf)
			readersize += n
			if n > 0 {
				out <- int(binary.BigEndian.Uint64(buf))  //数据压入
			}
			if err != nil || (chunksize != -1 && readersize >= chunksize) {
				break
			}
		}
		close(out)
	}()
	return out
}

func ArraySource(num...int) <- chan int {
	var out = make(chan int)
	go func() {
		for _,v := range num{
			out <- v //将数值数据压入
		}
		close(out)
	}()
	return out
}



