package main

import "fmt"

//a、找到第一个等于3的数  (只能顺序查找
// b找到最后一个等于3的
//c找到第一个大于等于2
//d找到最后一个小于7的
//
func Bin_searchD(arr []int,data int) int {
	left := 0
	right := len(arr)-1 //最上面最小面
	index := -1  //索引
	for left <= right {
		mid := (left+right)/2
		if arr[mid]>data{
			right = mid -1
		}else {
			if mid == len(arr)-1 || arr[mid+1]>data { //临界点
				index = mid
				break
			}else {
				left = mid+1
			}
		}
	}
	return index  //不存在
}

func Bin_searchC(arr []int,data int) int {
	left := 0
	right := len(arr)-1 //最上面最小面
	index := -1  //索引
	for left <= right {
		mid := (left+right)/2
		if arr[mid]<data{
			left = mid +1
		}else {
			if mid == 0 || arr[mid-1]<data { //临界点
				index = mid
				break
			}else {
				right = mid-1
			}
		}
	}
	return index  //不存在
}

func Bin_searchB(arr []int,data int) int {
	left := 0
	right := len(arr)-1 //最上面最小面
	index := -1  //索引
	for left <= right {
		mid := (left+right)/2
		if arr[mid]>data {
			right = mid - 1
		}else if arr[mid]<data{
			left = mid +1
		}else {
			if mid == len(arr)-1 || arr[mid+1]!=data{
				index = mid
				break
			}else {
				left = mid+1  //递归继续查找
			}
		}
	}
	return index  //不存在
}

//二分查找
func Bin_searchA(arr []int,data int) int {
	left := 0
	right := len(arr)-1 //最上面最小面
	index := -1  //索引
	for left <= right {
		mid := (left+right)/2
		if arr[mid]>data {
			right = mid - 1
		}else if arr[mid]<data{
			left = mid +1
		}else {
			if mid == 0 || arr[mid-1]!=data{
				index = mid
				break
			}else {
				right = mid-1  //递归继续查找
			}
		}
	}
	return index  //不存在
}

func main()  {
	arr := []int{1,2,3,3,3,3,3,4,5,6,6,6,6,7,9,10}
	fmt.Println(arr)
	fmt.Println(Bin_searchA(arr,3))  //查找第一个出现数据的位置
	fmt.Println(Bin_searchB(arr,3)) //查找最后一个出现数据的位置
	fmt.Println(Bin_searchC(arr,3))//第一个大于等于x的数据位置
	fmt.Println(Bin_searchD(arr,3))//最后一个小于等于x的数据位置
	/*
	for ; ;  {
		var inputdata int
		fmt.Println("输入要搜索的数据：")
		fmt.Scanf("%d",&inputdata)
		index := Bin_searchA(arr,inputdata)
		if index == -1 {
			fmt.Println("没有找到")
		}else {
			fmt.Println(index,arr[index])
		}
	}*/
}
