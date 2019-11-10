package main

//通过select进行任务调度
import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1500))*time.Millisecond)
			out <- i
			i ++
		}
	}()
	return out
}

func worker(id int,c chan int)  {
	for n:= range c{
		time.Sleep(time.Second)
		fmt.Printf("woker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {  //拿到这个chan的只能发数据
	c := make(chan int)
	go worker(id ,c)
	return c
}

func main() {
	var c1,c2 = generator() ,generator() //如果c1,c2都是空，select还能运行（进入default）相当于非阻塞
	var worker = createWorker(0)

	var values []int  //有一个缓存，避免存数据过快而消耗数据较慢导致数据跳点
	tm := time.After(10*time.Second)  //运行10秒后退出
	tick := time.Tick(time.Second) //每隔一段时间
	for  {
		var activeworker chan <- int
		var activevalue int

		if len(values)> 0{
			activeworker = worker
			activevalue = values[0]
		}
		select {
		case n := <- c1:
			values = append(values,n)
			//fmt.Println("recieved from c1:",n)
		case n := <- c2:
			values = append(values,n)
			//fmt.Println("recieved from c2:",n)
		case activeworker <- activevalue:
			values = values[1:]
		case <-time.After(800*time.Millisecond):
			fmt.Println("timeout") //每次select化的时间，每两次数据插入之间超过800毫秒（800毫秒内没生成数据
		//default:
			//fmt.Println("no value recieved")
		case <- tick:
			fmt.Println("queue len:",len(values))  //看队列长度
		case <- tm:  //从程序运行到现在总共的时间
			fmt.Println("bye!")
			return
		}
	}

}
