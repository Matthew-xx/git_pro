package main

import "fmt"

//快速排序（双冒泡排序）生活中最长用，速度最快
//5,3,4,2,6,7,1
//5  3,4,2,6,7,1  //  后边的与5比较
//3,4,2,1  5  6,7  //小于5的放一边，大于5的放一边
//3   4,2,1
//2,1  3   4

func QuickSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}else {
		splitdata := arr[0]  //以第一个为基准
		low := make([]int,0,0)  //存储比基准大的
		high := make([]int,0,0)  //存储比基准小的
		mid := make([]int,0,0)   //存储于基准相等的
		mid = append(mid,splitdata)   //首先加入基准数
		for i:=1;i<length;i++ {  //分块
			if arr[i] < splitdata {
				low = append(low,arr[i])
			}else if arr[i] > splitdata{
				high = append(high,arr[i])
			}else {
				mid = append(mid,arr[i])
			}
		}
		low,high = QuickSort(low),QuickSort(high)  //切割递归处理
		myarr := append(append(low,mid...),high...)
		return myarr
	}
}

//取最后一个数
func QuickSort2(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}else {
		n := length-1
		splitdata := arr[n]  //以第一个为基准
		low := make([]int,0,0)  //存储比基准大的
		high := make([]int,0,0)  //存储比基准小的
		mid := make([]int,0,0)   //存储于基准相等的
		mid = append(mid,splitdata)   //首先加入基准数
		for i:=0;i<length;i++ {  //分块
			if i == n {
				continue
			}
			if arr[i] < splitdata {
				low = append(low,arr[i])
			}else if arr[i] > splitdata{
				high = append(high,arr[i])
			}else {
				mid = append(mid,arr[i])
			}
		}
		low,high = QuickSort(low),QuickSort(high)  //切割递归处理
		myarr := append(append(low,mid...),high...)
		return myarr
	}
}


func main()  {
	arr := []int{3, 1, 4, 9, 6, 7, 5}
	fmt.Println(QuickSort(arr))
}
