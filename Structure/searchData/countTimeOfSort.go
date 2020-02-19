package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"github.com/pkg/errors"
	//"strings"
	"time"
)

//快速排序
func QuickSortT(arr []string) []string {
	length := len(arr)
	if length <= 1 {
		return arr
	}else {
		splitdata := arr[0]  //以第一个为基准
		low := make([]string,0,0)  //存储比基准大的
		high := make([]string,0,0)  //存储比基准小的
		mid := make([]string,0,0)   //存储于基准相等的
		mid = append(mid,splitdata)   //首先加入基准数
		for i:=1;i<length;i++ {  //分块
			if arr[i] < splitdata {
				low = append(low,arr[i])
			}else if arr[i] > splitdata{
				high = append(high,arr[i])
			}else {
				mid = append(mid,arr[i])
			}
		}
		low,high = QuickSortT(low),QuickSortT(high)  //切割递归处理
		myarr := append(append(low,mid...),high...)
		return myarr
	}
}

func main()  {
	t1 := time.Now()
	const N = 27  //需要开辟的内存
	allstrs := make([]string,N,N)  //开辟数组存储数据
	//fmt.Println(len(allstrs))

	fi,err := os.Open("F:\\Software\\go_path\\my_pro\\Structure\\data\\datasort.txt")
	if err != nil {
		errors.New("文件读取失败")
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	i := 0
	for {  //循环，逐行读取数据
		line,_,err := br.ReadLine() //读取一行
		if err == io.EOF {
			break  //跳出循环
		}
		//fmt.Println(string(line))
		linestr := string(line)  //转换为字符串
		//mystrs := strings.Split(linestr,"  #  ")  //字符串切割
		allstrs[i] = linestr
		i++
	}

	fmt.Println("读取完成")
	time.Sleep(3*time.Second)

	allstrs = QuickSortT(allstrs)
	fmt.Println(allstrs)
	fmt.Println("排序完成")
	used := time.Since(t1)
	fmt.Println(used)

	path := "F:\\Software\\go_path\\my_pro\\Structure\\data\\datasortTime.txt"  //保存文件的路径
	savefile,_ := os.Create(path)
	defer savefile.Close()
	savebuff := bufio.NewWriter(savefile)
	for i:=0;i<len(allstrs);i++ {
		fmt.Fprintln(savebuff,allstrs[i])
	}
	savebuff.Flush()
}


