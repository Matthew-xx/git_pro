package main

import (
	"fmt"
)

//冒泡排序(两两互换
//9,8,7,10,12,3,4
//8,9,7,10,12,3,4
//8,7,9,10,12,3,4
//8,7,9,10,3,12,4

func BubbleFindMax(arr []int) int {
	length := len(arr)
	if length <= 1 {
		return arr[0]
	}else {
		for i:=0;i<length-1;i++ {
			if arr[i] > arr[i+1] {
				arr[i],arr[i+1] = arr[i+1],arr[i]
			}
		}
		return arr[length-1] //最大值
	}
}

func BubbleSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}else {
		for i:=0;i<length-1;i++ {  //只剩最后一个，则不需要冒泡
			isneedexchange := false  //假定不需要交换，一交换就设定为真
			for j:=0;j<length-1-i;j++ {
				if arr[j] > arr[j+1] {
					arr[j],arr[j+1] = arr[j+1],arr[j]
					isneedexchange = true
				}
			}
			if !isneedexchange { //当
				break  //当始终false（不需要交换时表示已排好），则退出。
			}
			fmt.Println(arr)
		}
		return arr
	}
}

func main() {
	arr := []int{3, 1, 4, 9, 6, 7, 5}
	//fmt.Println(BubbleFindMax(arr))
	fmt.Println(BubbleSort(arr))
}


