package main

import (
	"fmt"
	"time"
)

//休眠排序，并发场景（多线程，分布式
var container chan bool
var flag bool
var count int

//休眠(根据指定休眠数据休眠时间输入管道，也就是小的休眠时间先排。也是一个排序
func tosleep(data int)  {
	time.Sleep(time.Duration(data)*time.Second)  //一般设置微秒
	fmt.Println("sleep:",data)
	container <- true  //管道输入OK
}

//监听
func listen(size int)  {
	for flag{ //flag为真时处理管道
		select {
		case <- container:
			count ++
			if count >= size{
				flag = false
				break   //等待5个数据采集完成就退出
			}
		}
	}
}
func main()  {
	var array []int = []int{17,2,8,25,31}
	flag = true  //标识可以写入（用于区分
	container = make(chan bool,5)  //5个管道
	for i:=0; i<len(array); i++ {
		go tosleep(array[i])  //因为是并发场合，所以可以插入i
	}
	go listen(len(array))

	for flag {
		time.Sleep(1*time.Second)
	}  //flag为false时结束
}

