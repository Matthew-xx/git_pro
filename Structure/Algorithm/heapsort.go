package main

import "fmt"

//堆排序

func HeapSortMax(arr []int,length int) []int {
	//length := len(arr)
	if length <= 1 {
		return arr
	}else {
		 depth := length/2-1   //深度(类似二叉树
		for i:=depth;i>=0;i-- {  //循环所有的三节点，n,2*n+1,2*n+2
			topmax := i  //假定最大的在i的位置
			leftchild := 2*i + 1  //左边
			rightchild := 2*i + 2 //右子节点
			if leftchild<= length-1 && arr[leftchild] > arr[topmax] { //作一判断防止越界
				topmax = leftchild  //如果左子节点较大，记录最大
			}
			if rightchild<= length-1 && arr[rightchild] > arr[topmax] {
				topmax = rightchild  //如果右子节点较大，记录最大
			}
			if topmax != i{ //确保I的值最大
				arr[i],arr[topmax] = arr[topmax],arr[i]
			}
		}
		return arr
	}
}

//每次先把一个最大的数（a[n]）放在最后，再从其他数中取出最大的数
func HeapSort(arr []int) []int {
	length := len(arr)
	for i:=0;i<length;i++ {
		lastmesslen := length -i //每次截取一段
		HeapSortMax(arr,lastmesslen)
		fmt.Println(arr)
		if i<length {
			arr[0],arr[lastmesslen-1] = arr[lastmesslen-1],arr[0]
		}
		fmt.Println("ex",arr)
	}
	return arr
}

func main()  {
	arr := []int{3, 1, 4, 9, 6, 7, 5}
	//fmt.Println(HeapSortMax(arr,2))
	fmt.Println(HeapSort(arr))
}