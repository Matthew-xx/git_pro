package main

import "fmt"

//快速排序与二分查找
func QuickSort(arr []int) []int {
	length := len(arr)
	if length <=1{
		return arr
	}else {
		splitdata := arr[0]  //第一个数字
		low := make([]int,0,0)
		high := make([]int,0,0)
		mid := make([]int,0,0)  //与相等的数字
		mid = append(mid,splitdata) //保存分离的数据
		//数据分3段处理，分别是大于、小于、等于基准数字
		for i:=1;i<length;i++ {
			if arr[i]<splitdata {
				low = append(low,arr[i])
			}else if arr[i]>splitdata{
				high = append(high,arr[i])
			}else {
				mid = append(mid,arr[i])
			}
		}
		low,high = QuickSort(low),QuickSort(high)  //递归循环
		myarr := append(append(low,mid...),high...)  //数据归并
		return myarr
	}
}

//二分查找
func Bin_search(arr []int,data int) int {
	left := 0
	right := len(arr)-1 //最上面最小面
	for left <= right {
		mid := (left+right)/2
		if arr[mid]>data {
			right = mid - 1
		}else if arr[mid]<data{
			left = mid +1
		}else {
			return mid  //代表找到
		}
	}
	return -1  //不存在
}

func main()  {
	arr := []int{2,4,1,13,9,7,10,5}
	fmt.Println(arr)
	fmt.Println(QuickSort(arr))

	var inputdata int
	fmt.Println("输入要搜索的数据：")
	fmt.Scanf("%d",&inputdata)
	index := Bin_search(arr,inputdata)
	if index == -1 {
		fmt.Println("没有找到")
	}else {
		fmt.Println(index,arr[index])
	}
}
