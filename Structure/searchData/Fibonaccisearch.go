package main

import (
	"fmt"
)

//斐波那契(数组不够时补足以凑够斐波那契数
func makeFibArray(arr []int) []int {
	length := len(arr)
	fiblen := 2
	first,second,third := 1,1,2
	for third<length{ //找出最接近的斐波那契
		//叠加计算斐波那契
		third,first,second = first+second,second,third
		fiblen++
	}
	fb := make([]int,fiblen)
	fb[0]=1
	fb[1]=1
	for i:=2; i<fiblen; i++ {
		fb[i] = fb[i-1]+fb[i-2]
	}
	return fb
}

//斐波那契搜索
func Fib_search(arr []int,val int) int {
	length := len(arr)
	fabarr := makeFibArray(arr)  //定制匹配的斐波那契数组
	filllength := fabarr[len(fabarr)-1]  //填充长度

	fillarr := make([]int,filllength)   //填充的数组
	for i,v := range arr{
		fillarr[i] = v
	}

	lastdata := arr[length-1] // 填充到最后一个数
	for i:=length;i<filllength;i++ {
		fillarr[i] = lastdata
	}

	left,mid,right := 0,0,length  //类似二分查找
	kindex := len(fabarr)-1  //起游标作用
	for left <=right{
		mid = left+fabarr[kindex-1]-1 //斐波那契切割
		if val < fillarr[mid] {
			right = mid -1 //类似二分查找
			kindex--
		}else if val > fillarr[mid]{
			left = mid + 1
			kindex-=2
		}else {
			if mid>right {  //越界
				return right
			}else {
				return mid
			}
		}
	}
	return -1  //不存在
}
func main()  {
	arr := make([]int,1000,1000)
	for i:=0; i<1000; i++ {
		arr[i]=i
	}

	for ; ;  {
		var inputdata int
		fmt.Println("输入要搜索的数据：")
		fmt.Scanf("%d",&inputdata)
		index := Fib_search(arr,inputdata)
		if index == -1 {
			fmt.Println("没有找到")
		}else {
			fmt.Println(index,arr[index])
		}
	}

}



