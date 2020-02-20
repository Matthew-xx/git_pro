package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

//快速排序与二分查找
func QuickSortStruct(arr []QQs) []QQs {
	length := len(arr)
	if length <=1{
		return arr
	}else {
		splitdata := arr[0].Quser  //第一个数字
		low := make([]QQs,0,0)
		high := make([]QQs,0,0)
		mid := make([]QQs,0,0)  //与相等的数字
		mid = append(mid,arr[0]) //保存分离的数据
		//数据分3段处理，分别是大于、小于、等于基准数字
		for i:=1;i<length;i++ {
			if arr[i].Quser<splitdata {
				low = append(low,arr[i])
			}else if arr[i].Quser>splitdata{
				high = append(high,arr[i])
			}else {
				mid = append(mid,arr[i])
			}
		}
		low,high = QuickSortStruct(low),QuickSortStruct(high)  //递归循环
		myarr := append(append(low,mid...),high...)  //数据归并
		return myarr
	}
}

//二分查找
func Bin_searchStruct(arr []QQs,data int) int {
	left := 0
	right := len(arr)-1 //最上面最小面
	for left < right {
		mid := (left+right)/2
		if arr[mid].Quser>data {
			right = mid - 1
		}else if arr[mid].Quser<data{
			left = mid +1
		}else {
			return mid  //代表找到
		}
	}
	return -1  //不存在
}

type QQs struct {
	Quser   int
	Qpass   string
	Qname   string
}

func main()  {

	const N = 27  //需要开辟的内存
	allstrs := make([]QQs,N,N)  //开辟数组存储数据
	path := "F:\\Software\\go_path\\my_pro\\Structure\\data\\data.txt"
	file, _ := os.Open(path)
	defer file.Close()

	br := bufio.NewReader(file)
	i := 0  //统计共多少行
	for  {
		line,_,end := br.ReadLine()  //逐行读取
		if end == io.EOF { //文件关闭
			break
		}

		linestr := string(line)
		lines := strings.Split(linestr,"  #  ")
		if len(lines) == 3 {
			allstrs[i].Quser,_ = strconv.Atoi(lines[0])
			allstrs[i].Qpass = lines[1]
			allstrs[i].Qname = lines[2]
			i++
		}
	}
	sortstartTime := time.Now()
	allstrs = QuickSortStruct(allstrs)
	fmt.Println("排序完成")
	fmt.Println("排序耗时：",time.Since(sortstartTime))
	fmt.Println("要搜索的数据：")
	var QQ int
	fmt.Scanf("%d",&QQ)  //查询QQ

	startTime := time.Now()
	index := Bin_searchStruct(allstrs,QQ)
	if index == -1{
		fmt.Println("查找不到数据")
	}else {
		fmt.Println("找到数据：",index,allstrs[index])
	}
	fmt.Println("搜索耗时：",time.Since(startTime))
}

