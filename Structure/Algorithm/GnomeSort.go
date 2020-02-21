package main

import "fmt"

//冒泡和插入的结合(侏儒排序：适合处理快排完了的数据
//7 0 8 2 16 5
//0 7 8 2 16 5
//0 7 2 8 16 5
//0 2 7 8 16 5   i--步骤


func GnomeSort(arr []int) []int {
	i := 1
	for ; i<len(arr); {
		if arr[i]>= arr[i-1] {
			i++  //符合顺序，继续前进
		}else {
			arr[i],arr[i-1] = arr[i-1],arr[i]
			if i>1 {
				i--  //完成一步交换再退回检查
			}
			fmt.Println(arr)
		}
	}
	return arr
}

func main()  {
	arr := []int{11,2,3,23,33,3,13,4,15,6,6,61,6,17,9,10}
	GnomeSort(arr)
	fmt.Println(arr)



}


