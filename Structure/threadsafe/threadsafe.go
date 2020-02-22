package main

import (
	"fmt"
	"time"
)

//线程安全:多个线程访问同个资源会产生资源竞争，会使得最终结果不正确

var money int

func add(pint *int)  {
	for i:=0;i<10000;i++ {
		*pint++
	}
}

func main()  {
	for i:=0;i<1000;i++ {
		go add(&money)
	}
	time.Sleep(time.Second*20)
	fmt.Println(money)
}

