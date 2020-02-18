package main
//选择排序(整数
import (
	"fmt"
)

//选择最大
func SelectSortMax(arr[] int) int {
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

//选择排序
func SelectSort(arr[] int) []int {
	length := len(arr)  //数组长度
	if length <=1{
		return arr  //一个元素的数组，直接返回
	}else {
		for i:=0;i<length-1;i++ { //只剩一个元素时不需挑选
			min := i   //标记索引
			for j:=i+1;j<length;j++ { //每次选出一个极小值
				if arr[min] < arr[j] {  //若是">"则是从小到大排序
					min = j //保存极大值的索引
				}
			}
			if i != min {
				arr[i],arr[min] = arr[min],arr[i]  //数据交换
			}
			fmt.Println(arr)
		}
		return arr
	}
}


func main()  {
	arr := []int{1,3,4,2,6,7,5}
	//fmt.Println(SelectSortMax(arr))
	fmt.Println(SelectSort(arr))
}
