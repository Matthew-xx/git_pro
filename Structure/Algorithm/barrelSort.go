package main

import "fmt"

//木桶排序
//超大型数据在范围有限的情况下，木桶比快速更快
//如果全国每个身高的人数(每个身高作一个桶，装着人数

//处理固定情况
func BucketSort(arr []int) []int {
	length := len(arr)
	if length <= 1{
		return arr
	}else {
		num := 4  //对要排的数据有一定了解，大概在什么范围
		buckets := make([][]int,num)  //创造二维数组
		for i:=0; i<length; i++ {
			buckets[arr[i]-1] = append(buckets[arr[i]-1],arr[i])   //木桶计数加1，arr[i]-1因为索引是从0开始
		}
		fmt.Println(buckets)
		tmppose := 0  //木桶排序
		for i:=0; i<num; i++ {
			bucketslen := len(buckets[i])  //求某一段的长度
			if bucketslen >0{
				copy(arr[tmppose:],buckets[i])   //展开数据
				tmppose += bucketslen  //追加长度，起定位作用
			}
		}
		return arr
	}
}

//处理不固定情况
func BucketSortX(arr []int) []int {
	length := len(arr)
	if length <= 1{
		return arr
	}else {
		num := length  //对要排的数据有一定了解，大概在什么范围
		max := SelectSortM(arr)
		index := 0  //默认的索引
		buckets := make([][]int,num)  //创造二维数组
		for i:=0; i<length; i++ {
			index = arr[i]*(num-1)/max  //求出的大概值，木桶的自动分配算法
			buckets[index] = append(buckets[index],arr[i])   //木桶计数加1，arr[i]-1因为索引是从0开始
		}
		fmt.Println(buckets)
		tmppose := 0  //木桶排序
		for i:=0; i<num; i++ {
			bucketslen := len(buckets[i])  //求某一段的长度
			if bucketslen >0{
				buckets[i] = SelectSortB(buckets[i]) //对木桶内部数据排序
				copy(arr[tmppose:],buckets[i])   //展开数据
				tmppose += bucketslen  //追加长度，起定位作用
			}
		}
		return arr
	}
}
//查找数组中最大值
func SelectSortM(arr []int) int {
	length := len(arr)
	if length <= 1{
		return arr[0]
	}else {
		max := arr[0]
		for i:=1; i<length; i++ {
			if arr[i]>max {
				max = arr[i]
			}
		}
		return max
	}
}

//选择排序
func SelectSortB(arr[] int) []int {
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
	arr := []int{1,2,3,4,4,1,3,2,2,3,1}
	fmt.Println(arr)

	fmt.Println(BucketSortX(arr))
}

