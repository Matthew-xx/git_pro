package main

import "fmt"

//梳子排序，是希尔排序的改进版
func CombSort(arr []int) []int {
	length := len(arr)
	gap := length
	for gap >1 {
		gap = gap*10/13
		for i:=0;i+gap<length;i++ {  //收缩
			if arr[i] > arr[gap+i] {
				arr[i],arr[gap+i] = arr[gap+i],arr[i]
			}
		}
	}
	return arr
}


func main() {
	var array []int = []int{17, 2, 8, 25, 31}
	fmt.Println(array)
	CombSort(array)
	fmt.Println(array)
}
