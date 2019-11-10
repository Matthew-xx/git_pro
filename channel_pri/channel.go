package main

import (
	"fmt"
	"time"
)
/*
func worker(id int,c chan int)  {
	for{
		fmt.Printf("woker %d received %c\n",id,<-c) //只能收数据
	}
}*/
//修改后
func worker(id int,c chan int)  {
	/*for{
		n, ok := <-c
		if !ok {
			break
		}
			fmt.Printf("woker %d received %c\n", id, n)
	}*/
	for n:= range c{
		fmt.Printf("woker %d received %c\n", id, n)
	}
}

func createWorker(id int) chan<- int {  //拿到这个chan的只能发数据
	c := make(chan int)
	go worker(id ,c)
	return c
}

func chanDemo()  {
	var channels [10]chan<- int
	for i :=0;i<10;i++{
		channels[i] = createWorker(i)
	}
	for i:=0; i<10; i++ {
		channels[i] <- 'a'+i
	}
	for i:=0; i<10; i++ {
		channels[i] <- 'A'+i
	}
}

func bufferedChannel()  { //缓冲通道
	c := make(chan int,3) //设置缓冲，在缓冲容量内只发不收都可以
	go worker(0,c)
	c <- 'a'  //有人发数据 就要有人收数据 ，
	c <- 'b'
	time.Sleep(time.Millisecond)
}

func channelClose()  { //数据收发完毕通过关闭，发送方的close，close后通知接收方，接收方再做判断（n,ok或range)
	c := make(chan int) //有无缓冲都可
	go worker(0,c)
	c <- 'a'  //有人发数据 就要有人收数据 ，
	c <- 'b'
	close(c)  //worker里面是循环接收，在下一句time.sleep间会收到很多空值，需对worker修改
	time.Sleep(time.Millisecond) //sleep以便进程打印完（但打印是在goroutine里面，并不知道他什么时候打印完，即不能无限休眠下去）改动见done
}

func main() {
	fmt.Println("1、channel as first-class citizen(一等公民，作为参数传)")
	chanDemo()
	fmt.Println("2、buffered channel(缓冲)")
	bufferedChannel()
	fmt.Println("3、channel close and range(关闭通道及range)")
	channelClose()
}
