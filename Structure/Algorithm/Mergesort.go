package main

import "fmt"

//归并排序(
//3, 1, 4, 9, 6, 7, 5
//3, 1, 4,    9, 6, 7, 5
//3, 1, 4,     9, 6,     7, 5   //分块排序
//1,3,4,  6,9   5,7
//

func InsertSortX(arr []int) []int {
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

//合并
func merge(leftarr []int,rightarr []int) []int {
	leftindex := 0
	rightindex := 0  //左右两边的索引
	lastarr := []int{}  //最终的数组
	for leftindex <len(leftarr)  && rightindex < len(rightarr) {
		//如果两边都没结束，需对两边进行循环（筛选
		if leftarr[leftindex] < rightarr[rightindex] {
			lastarr = append(lastarr,leftarr[leftindex])
			leftindex++
		}else if leftarr[leftindex] > rightarr[rightindex] {
			lastarr = append(lastarr,rightarr[rightindex])
			rightindex++
		}else {
			lastarr = append(lastarr,leftarr[leftindex])
			lastarr = append(lastarr,rightarr[rightindex])
			rightindex++
			leftindex++
		}
	}
	for leftindex < len(leftarr){  //把没有结束的块归并过来
		lastarr = append(lastarr,leftarr[leftindex])
		leftindex++
	}
	for rightindex < len(rightarr){
		lastarr = append(lastarr,rightarr[rightindex])
		rightindex++
	}

	return lastarr
}

func MergeSort(arr []int) []int  {
	length := len(arr)
	if length <= 1 {
		return arr  //改进，可在length小于N时调用其他排序算法
	}else if  length >1 && length <5{
		return InsertSortX(arr)
	}else {
		mid := length / 2
		leftarr := MergeSort(arr[:mid])  //循环递归（左边部分
		rightarr := MergeSort(arr[mid:])  //循环递归（右边部分

		return merge(leftarr,rightarr)
	}
}
func main()  {
	arr := []int{3, 1, 4, 9, 6, 7, 5,8}
	fmt.Println(MergeSort(arr))
}