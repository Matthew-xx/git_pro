package main

import "fmt"

//三分查找

func ThirdSearch(arr []int,data int) int {
	low := 0
	high := len(arr) -1  //确定底部与高部
	i := 0
	for low <= high{
		mid1 := low + int((high-low)/3)
		mid2 := high - int((high-low)/3)

		i++
		fmt.Println("次数：",i)

		middata1 := arr[mid1]
		middata2 := arr[mid2]
		if middata1 == data{
			return mid1
		}else if middata2 == data{
			return mid2
		}

		if middata1 < data {
			low = mid1 + 1
		}else if middata2 > data {
			high = mid2 - 1
		}else {
			low = low + 1
			high = high - 1
		}
	}
	return -1
}


func main()  {
	arr := make([]int,1000,1000)
	for i:=0;i<1000;i++ {
		arr[i] = i
	}
	fmt.Println(ThirdSearch(arr,111))
}
