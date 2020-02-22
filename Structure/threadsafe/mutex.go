package main

import (
	"fmt"
	"sync"
	"time"
)


//解决线程安全

var money1 int
var lock *sync.RWMutex=new(sync.RWMutex)  //new初始化

func add1(pint *int)  {
	lock.Lock()  //锁
	for i:=0;i<10000;i++ {
		*pint++
	}
	lock.Unlock()  //解锁
}

func main()  {
	for i:=0;i<1000;i++ {
		go add1(&money1)
	}
	time.Sleep(time.Second*20)
	fmt.Println(money1)
}

