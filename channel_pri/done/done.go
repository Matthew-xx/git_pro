package main

import (
	"fmt"
)

func doWorker(id int,c chan int, done chan bool)  {

	for n:= range c{
		fmt.Printf("woker %d received %c\n", id, n)
		go func() {done <- true}() //通知外面事情做完了.并行，这样就不至于下面两个for循环发数据阻塞
	}
}

type worker struct {
	in chan int
	done chan bool
}  //构造变量以便传入

func createWorker(id int) worker {  //拿到这个chan的只能发数据
	w := worker{  //传入变量
		in : make(chan int),
		done: make(chan bool),
	}
	go doWorker(id ,w.in,w.done)
	return w
}

func chanDemo()  {
	var wokers [10]worker
	for i :=0;i<10;i++{
		wokers[i] = createWorker(i)
	}
	for i,worker := range wokers{
		worker.in <- 'a'+i
		//<- wokers[i].done  //接收结束，但这样会是顺序执行(发一条数据等待一条数据接收完毕才执行下一条
	}
	for i,worker := range wokers{
		worker.in <- 'A'+i //阻塞的
		//<- wokers[i].done
	}
	//全部结束后才退出来全部发完才接收全部的done
	for _,worker := range wokers{
		<- worker.done
		<- worker.done
	}
}

func main() {
	fmt.Println("1、channel as first-class citizen(一等公民，作为参数传)")
	chanDemo()

}
