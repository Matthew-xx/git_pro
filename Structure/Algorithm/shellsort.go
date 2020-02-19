package main

import "fmt"

//希尔排序(步长收缩）多用在并发场合
//3, 1, 4, 9, 6, 7, 5,8
//3           6
//   1           7
//       4          5
//          9          8   //根据这两个数的大小改变顺序
//3  1  4   8  6 7  5  9
//1      4     6    5      //改变步长，同上

func ShellSortStep(arr []int,start,gap int)  {
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

func ShellSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr  //改进，可在length小于N时调用其他排序算法
	}else{
		gap := length / 2   //设置步长
		for gap > 0 {
			for i:=0;i<gap;i++ {
				ShellSortStep(arr,i,gap)
			}
			gap/=2   //gap--
		}
	}
	return arr
}

func main()  {
	arr := []int{3, 1, 4, 9, 6, 7, 5,8}
	fmt.Println(ShellSort(arr))
}