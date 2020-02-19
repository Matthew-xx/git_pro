package main

import (
	"fmt"
)

//散列桶
//固定的：如银行用户的存款（以百万，千万，亿分别作桶里面排序
//非固定的：如800万考生高考分数、1亿人身高排序


func SelectSortMax1(arr[] int) int {
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


//按位数进行排序
func BitSort(arr []int,bit int) []int {
	length := len(arr)
	bitcounts := make([]int,10)
	for i:=0;i<length;i++ {
		num := (arr[i] / bit) % 10  //取余数，分层处理，bit=1001时，三位数就不参与排序
		bitcounts[num]++   //统计余数相等的个数

	}
	fmt.Println(bitcounts)
	for i:=1;i<10;i++ {//累加
		bitcounts[i] += bitcounts[i-1]  //叠加后可计算位置
	}
	fmt.Println(bitcounts)

	tmp := make([]int,10)
	//计算位置
	for i:=length-1;i>=0;i-- {
		num := (arr[i] / bit) % 10  //分层处理，bit=1001时，三位数就不参与排序
		tmp[bitcounts[num]-1] = arr[i]  //计算排序的位置
		bitcounts[num]--
	}
	for i:=0;i<length;i++ {
		arr[i] = tmp[i]  //保存数组
	}
	return arr
}

func RadixSort(arr []int) []int {
	max := SelectSortMax1(arr)  //,寻找数组最大值,需根据极大值确定数量级
	for bit:=1;max/bit>0;bit*=10 {  //当max/bit<0,说明已经小于max
		//按数量级分段
		arr = BitSort(arr,bit)  //每次处理一个级别的排序（个位数、十位数、、
		fmt.Println(arr)
	}
	return arr
}

func main()  {
	arr := []int{150,38,41,204,63,72,507,4123,4447,118}
	fmt.Println(RadixSort(arr))
}
