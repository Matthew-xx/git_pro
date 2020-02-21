package main

import "fmt"

//斐波那契搜索(中值搜索:只需搜索一次。配合改进版的快速排序，速度能提升10倍）

//二分查找
func Mid_search(arr []int,data int) int {
	left := 0
	right := len(arr)-1 //最上面最小面
	i := 0
	for left < right {
		i++
		fmt.Printf("第%d次搜索\n",i)

		leftv := float64(data-arr[left])  //左边到中段的数据
		allv := float64(arr[right]-arr[left])  //整段数据
		diff := float64(right-left)  //数据长度中间值
		mid := int(float64(left) + leftv/allv*diff)   //计算中间值

		if mid <0 || mid >=len(arr){
			return -1
		}
		//在搜索数据的时候，使用下面的mid做中间值时，搜索4需要8次(10以上的数据要更久)。而上面的mid只需一次
		//mid := (left+right)/2
		if arr[mid]>data {
			right = mid - 1
		}else if arr[mid]<data{
			left = mid +1
		}else {
			return mid  //代表找到
		}
	}
	return -1  //不存在
}
func main()  {
	arr := make([]int,1000,1000)
	for i:=0; i<1000; i++ {
		arr[i]=i
	}

	var inputdata int
	fmt.Println("输入要搜索的数据：")
	fmt.Scanf("%d",&inputdata)
	index := Mid_search(arr,inputdata)
	if index == -1 {
		fmt.Println("没有找到")
	}else {
		fmt.Println(index,arr[index])
	}
}


