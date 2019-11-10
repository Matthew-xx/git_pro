package main

import (
	"fmt"
	"sync"
)

func doWork(id int,w work)  {

	for n:= range w.in{
		fmt.Printf("woker %d received %c\n", id, n)
		w.done()
	}
}

type work struct {
	in chan int
	done func()
}  //构造变量以便传入

func createWork(id int, wg *sync.WaitGroup) work {
	w := work{  //传入变量
		in : make(chan int),
		done: func() {wg.Done()},
	}
	go doWork(id ,w)
	return w
}

func chandemo()  {
	var wg sync.WaitGroup  //使用传统的同步机制
	var works [10]work
	for i :=0;i<10;i++{
		works[i] = createWork(i,&wg)
	}

	wg.Add(20)  //两个10共20
	for i,work := range works{
		work.in <- 'a'+i
		//wg.Add(1) //也可以在循环里面添加
	}
	for i,work := range works{
		work.in <- 'A'+i //阻塞的
	}
	//全部结束后才退出来全部发完才接收全部的done
	wg.Wait()
}

func main() {
	fmt.Println("1、channel as first-class citizen(一等公民，作为参数传)")
	chandemo()

}
