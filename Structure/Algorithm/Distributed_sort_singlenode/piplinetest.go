package main

import (
	"./pipelineMiddleware"
	"bufio"
	"fmt"
	"os"
	"time"
)

func main1()  {
	var filename = "data.in"  //文件写入
	var count = 100000
	file,_ := os.Create(filename)
	defer file.Close()

	mypipe := pipelineMiddleware.RandSource(count)  //管道装随机数
	writer := bufio.NewWriter(file)  //xr
	pipelineMiddleware.WriteSlink(writer,mypipe) //将mypipe写入
	writer.Flush() //刷新

	file,_ = os.Open(filename)
	defer file.Close()
	mypipex := pipelineMiddleware.ReaderSource(bufio.NewReader(file),-1)
	countx := 0
	for v:= range  mypipex{
		fmt.Println(v)
		countx++
		if countx > 1000{
			break
		}
	}
}

func main()  {
	go func() {  //调用多线程
		myp := pipelineMiddleware.Merge(
			pipelineMiddleware.InMemorySort(pipelineMiddleware.ArraySource(3,9,4,2,1,7)),
			pipelineMiddleware.InMemorySort(pipelineMiddleware.ArraySource(13,19,14,12,11,27)),
		)
		for v:= range myp{
			fmt.Println(v)
		}
	}()
	//如果没有调用go func 会出错（系统认为死锁）：
	// fatal error: all goroutines are asleep - deadlock!
	time.Sleep(time.Second*2)
}



