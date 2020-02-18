package main

import "fmt"

//插入排序

//// 3 5 4 1 6 7
//// 3   5 4 1 6 7
//// 3 5    4 1 6 7
//// 3 4 5    1 6 7

func InsertTest(arr []int) []int {
	backup := arr[2]
	j := 2-1  //从上一个位置循环找到位置插入
	for j>=0 && backup<arr[j] {
		arr[j+1] = arr[j]
		j--
	}
	arr[j+1] = backup
	return arr
	
}

func InsertSort(arr []int) []int {
	length := len(arr)  //数组长度
	if length <=1{
		return arr  //一个元素的数组，直接返回
	}else {
		for i:=1;i<length;i++ {
			backup := arr[i]  //备份插入的数据
			j := i-1  //从上一个位置循环找到位置插入
			for j>=0 && backup<arr[j] {
				arr[j+1] = arr[j]  //从前往后移
				j--
			}
			arr[j+1] = backup  //插入
		}
		return arr
	}
}

func main()  {
	arr := []int{1,3,9,4,6,7,5,8,2}
	fmt.Println(InsertSort(arr))
}
