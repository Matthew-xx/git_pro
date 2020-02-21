package main

import (
	"fmt"
)
//找第K大可以用堆排序，也可用快速排序改进版
//快速排序
func QuickSortK(arr []int)  {
	quickSortGo(arr,0,len(arr)-1)
}

func quickSortGo(arr []int,left,right int)  {
	if left >= right {
		return
	}
	q := partition(arr,left,right)  //返回的切段的数据
	quickSortGo(arr,left,q-1)  //处理q-1(前段的数据）
	quickSortGo(arr,q+1,right) //处理q+1(后段的数据）
}

//分段
func partition(arr []int,left,right int) int {
	pivot := right
	i := left
	for j:=left;j<pivot;j++ {
		if arr[j] > arr[pivot] {  //求极大值
			swap(arr,i,j)
			i++
		}
	}
	swap(arr,i,pivot)
	return i
}

func swap(arr []int,i,j int)  {
	arr[i],arr[j] = arr[j],arr[i]
}

func findKlargest(arr []int,k int) int {
	return findKlargestgo(arr,0,len(arr)-1,k) //查找第k大的数
}

func findKlargestgo(arr []int,left,right int,k int) int {
	if left >= right{ //左右重合
		return arr[left]  //只剩一个元素，最大的元素
	}
	query := partition(arr,left,right)
	if query+1 == k{
		return arr[query]  //第k大的数
	}
	if k<(query+1) {
		return findKlargestgo(arr,left,query-1,k)  //再到限定字段递归查找,直到
	}

	return findKlargestgo(arr,left,query+1,k)
}

func main()  {
	arr := []int{11,2,3,23,33,3,13,4,15,6,6,61,6,17,9,10}
	fmt.Println(findKlargest(arr,10))
	fmt.Println(arr)
	QuickSortK(arr)
	fmt.Println(arr)
}




