package main

import "fmt"

//计数排序


func SelectSortMaxy(arr[] int) int {
	length := len(arr)  //数组长度
	if length <=1{
		return arr[0]  //一个元素的数组，直接返回
	}else {
		max := arr[0]  //假定第一个最大
		for i:=1;i<length;i++ {
			if arr[i] > max { //任何一个较大的数，则替换最大
				max = arr[i]
			}
		}
		return max
	}
}

func CountSort(arr []int) []int {
	max := SelectSortMaxy(arr)

	sortedarr := make([]int,len(arr))  //排序之后存储

	countarr := make([]int,len(arr))  //统计次数
	for _,v := range arr{
		countarr[v]++
	}
	fmt.Println("第一次统计次数：",countarr)

	for i:=1;i<=max;i++ {
		countarr[i] += countarr[i-1] //叠加
	}
	fmt.Println("次数叠加：",countarr)

	for _,v := range arr{
		sortedarr[countarr[v]-1] = v //展开数据,countarr[v]-1按照次数的计算的位置

		countarr[v]--
		fmt.Println("zkcount",countarr)
		fmt.Println("zk",sortedarr)
	}
	return sortedarr
}
func main()  {
	arr := []int{1,2,5,1,4,1,2,4,1,5,2,2,3,7,6,7,5,6}
	fmt.Println(CountSort(arr))
}



