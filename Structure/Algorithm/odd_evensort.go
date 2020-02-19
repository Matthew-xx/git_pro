package main

import "fmt"

//奇偶排序,只在一种场合用得多(方差不是很大的情况下
//对arr[奇]排序，对arr[偶]排序

func OddEven(arr []int) []int {
	isSorted := false  //奇偶不需要交换的时候

	for ;isSorted==false; {
		isSorted = true
		for i:=1;i<len(arr)-1;i=i+2 { //奇数位
			if arr[i] > arr[i+1] {
				arr[i],arr[i+1] = arr[i+1],arr[i]
				isSorted = false
			}
		}
		fmt.Println("1",arr)
		for i:=0;i<len(arr)-1;i=i+2 {//偶数位
			if arr[i] > arr[i+1] {
				arr[i],arr[i+1] = arr[i+1],arr[i]
				isSorted = false
			}
		}
		fmt.Println("0",arr)
	}
	return arr
}

func main()  {
	arr := []int{3, 1, 4, 9, 6, 7, 5}
	fmt.Println(OddEven(arr))
}