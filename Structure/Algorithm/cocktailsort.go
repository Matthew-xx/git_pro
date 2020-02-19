package main

import "fmt"

//鸡尾酒排序：双向冒泡（正向及反向冒泡，得到一个最大及一个最小值

func CocktailSort(arr []int) []int {
	for i:=0;i<len(arr)/2;i++ {  //因为是双向冒泡，所以次数相当于减半
		left := 0
		right := len(arr)-1 //左右两边的位置
		for left <= right {  //结束的条件
			if arr[left] > arr[left+1] {
				arr[left],arr[left+1] = arr[left+1],arr[left]
			}
			left++   //从左往右冒泡

			if arr[right-1] > arr[right] {
				arr[right-1],arr[right] = arr[right],arr[right-1]
			}
			right--    //从右往左冒泡
		}
		fmt.Println(arr)
	}
	return arr
}


func main()  {
	arr := []int{1,3,4,8,2,6,7,5}
	fmt.Println(CocktailSort(arr))
}


