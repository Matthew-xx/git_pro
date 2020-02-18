package main

import (
	"fmt"
	"strings"
)

//选择最大
func SelectSortMaxString(arr[] string) string {
	length := len(arr)  //数组长度
	if length <=1{
		return arr[0]  //一个元素的数组，直接返回
	}else {
		max := arr[0]  //假定第一个最大
		for i:=1;i<length;i++ {
			if strings.Compare(arr[i],max) > 0 {
				max = arr[i]
			}
			/*
			if arr[i] > max { //任何一个较大的数，则替换最大
				max = arr[i]
			}*/
		}
		return max
	}
}

//选择排序
func SelectSortString(arr[] string) []string {
	length := len(arr)  //数组长度
	if length <=1{
		return arr  //一个元素的数组，直接返回
	}else {
		for i:=0;i<length-1;i++ { //只剩一个元素时不需挑选
			min := i   //标记索引
			for j:=i+1;j<length;j++ { //每次选出一个极小值
			/*
				if arr[min] < arr[j] {  //若是">"则是从小到大排序
					min = j //保存极大值的索引
				}*/
				if strings.Compare(arr[min],arr[j]) < 0 {
					min = j
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

func main1()  {  //字符串的两种比较
	fmt.Println("a"<"b")  //字符串存在地址的比较(用处较小
	fmt.Println(strings.Compare("a","b")) //小于返回 -1
	fmt.Println(strings.Compare("g","b"))
	fmt.Println(strings.Compare("ac","au"))
}

func main()  {
	arr := []string{"d","f","u","n","w","o"}
	//fmt.Println(SelectSortMaxString(arr))
	fmt.Println(SelectSortString(arr))
}


