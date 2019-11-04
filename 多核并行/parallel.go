package main

import (
	"fmt"
	"runtime"
	"time"
)
//计算任务：启动多个子协程，子协程数量和 CPU 核心数保持一致，以便充分利用多核并行运算，每个子协程计算分给它的那部分计算任务，
// 最后将不同子协程的计算结果再做一次累加，这样就可以得到所有数据的计算总和

func sum(seq int, ch chan int) {
	defer close(ch)
	sum := 0
	for i := 1; i <= 10000000; i++ {
		sum += i
	}
	fmt.Printf("子协程%d运算结果:%d\n", seq, sum)
	ch <- sum
}

func main()  {
	// 启动时间
	start := time.Now()
	// 最大 CPU 核心数
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
	chs := make([]chan int, cpus)
	for i := 0; i < len(chs); i++ {
		chs[i] = make(chan int, 1)
		go sum(i, chs[i])
	}
	sum := 0
	for _, ch := range chs {
		res := <- ch
		sum += res
	}
	// 结束时间
	end := time.Now()
	// 打印耗时
	fmt.Printf("最终运算结果: %d, 执行耗时(s): %f\n", sum, end.Sub(start).Seconds())
}
