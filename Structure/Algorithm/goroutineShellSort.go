package main

//多线程希尔排序
import (
	"fmt"
	"runtime"
	"sync"
)

//希尔排序(步长收缩）多用在并发场合
//3, 1, 4, 9, 6, 7, 5,8
//3           6
//   1           7
//       4          5
//          9          8   //根据这两个数的大小改变顺序
//3  1  4   8  6 7  5  9
//1      4     6    5      //改变步长，同上

func ShellSortGoRoutine(arr []int)  {
	if len(arr)<2 || arr == nil {
		return  //数组为空或数组只有一个元素无需排序
	}
	wg :=sync.WaitGroup{}  //等待多个线程返回
	GoRoutinenum := runtime.NumCPU()  //抓取系统的CPU个数
	//压缩空间
	for gap:=len(arr)/2; gap>0; gap/=2 {
		wg.Add(GoRoutinenum)
		ch := make(chan int ,10000)  //通道，进行线程通信
		go func() {
			//管道写入任务
			for k:=0; k<gap; k++ {
				ch <- k
			}
			close(ch)  //关闭管道
		}()
		for k:=0;k<GoRoutinenum;k++ {
			go func() {
				for v:=range ch{
					ShellSortStepGR(arr,v,gap)  //完成一个步骤的排序
				}
				wg.Done()  //等待完成
			}()
		}
		wg.Wait() //等待
	}
	fmt.Println(arr)
}

func ShellSortStepGR(arr []int,start,gap int)  {
	length := len(arr)

	for i:=start+gap;i<length;i+=gap { //插入排序的变种
		backup := arr[i]  //备份插入的数据
		j := i-gap  //从上一个位置循环找到位置插入
		for j>=0 && backup<arr[j] {
			arr[j+gap] = arr[j]  //从前往后移
			j-=gap
		}
		arr[j+gap] = backup  //插入
	}
}

func ShellSortGR(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr  //改进，可在length小于N时调用其他排序算法
	}else{
		gap := length / 2   //设置步长
		for gap > 0 {
			for i:=0;i<gap;i++ {
				ShellSortStepGR(arr,i,gap)
			}
			gap/=2   //gap--
		}
	}
	return arr
}

func main()  {
	arr := []int{3, 1, 4, 9, 6, 7, 5,8}
	ShellSortGoRoutine(arr)

}
