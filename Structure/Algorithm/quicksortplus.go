package main

//改进版：始终在一个数组内交换数据，避免一直开新的数组空间。省内存

import (
	"fmt"
	"math/rand"
)

func SortForMerge(arr []int,left,right int)  {
	for i:=left; i<=right; i++{
		temp := arr[i]  //备份数据
		var j int
		for j=i; j>left && arr[j-1]>temp; j-- { //定位
			arr[j] = arr[j-1]  //数据往后移动
		}
		arr[j] = temp //插入
	}
}

//数据交换
func swap(arr []int,i,j int)  {
	arr[i],arr[j] = arr[j],arr[i]
}

//递归快速排序
//left,right标识下位置
func QuickSortX(arr []int,left,right int)  {
	if right-left<3 { //数组剩下2个数时直接插入排序（一般小于15都可插入
		SortForMerge(arr,left,right)
	}else { //快速排序
		//随机定一个数字放第一个位置，
		swap(arr,left,rand.Int()%(right-left+1)+left)
		vdata := arr[left]  //坐标数据，小于的放左边，大于该数字的放右边
		//分三部分
		lt := left  // 使得 arr[left+1,...,lt] < vdata
		rt := right+1 //使得arr[rt,...,right] > vdata
		i := left +1  //使得arr[lt+1,...,i] == vdata
		for i < rt{
			if arr[i] < vdata { //往左边移动
				swap(arr,i,lt+1) //移动到小于的地方
				lt++   //前进循环
				i++
			}else if arr[i] > vdata { //往右边移动
				swap(arr,i,rt-1)  //移动到大于的地方
				rt--
			}else {
				i++
			}
		}
		swap(arr,left,lt)  //交换头部位置
		//再递归处理大于和小于的部分
		QuickSortX(arr,left,lt-1)  //相当于上步处理完后，再处理arr[i] < vdata这一部分
		QuickSortX(arr,rt,right)
	}
}


//快排改良的核心
func QuickSortPlus(arr []int){
	QuickSortX(arr,0,len(arr)-1)
}


func main()  {
	arr := []int{2,4,1,13,9,7,10,5,11,3}
	fmt.Println("排序前:",arr)
	QuickSortPlus(arr)
	fmt.Println("排序后:",arr)
}

