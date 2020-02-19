package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

//内存搜索,比硬盘搜索快

func main() {
	startTime := time.Now()
	const N = 27  //需要开辟的内存
	allstrs := make([]string,N,N)  //开辟数组存储数据
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
		allstrs[i] = string(line)
		i++
	}
	for j:=0;j<N;j++ {
		if strings.Contains(allstrs[j],"星") { //搜索字符串  “h"
			fmt.Println(allstrs[j])
		}
	}
	fmt.Println("搜索耗时：",time.Since(startTime))
	fmt.Printf("共发现 %d 行数据",i)
}

