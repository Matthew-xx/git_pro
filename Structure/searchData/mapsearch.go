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


func main()  {

	const N = 27  //需要开辟的内存
	allstrs := make(map[int]string,N)  //初始化map
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
			Quser,_ := strconv.Atoi(lines[0])
			Qpass := lines[1]
			//Qname := lines[2]
			allstrs[Quser] = Qpass  //映射到map中
			i++
		}
	}
	for ; ;  {
		fmt.Println("要搜索的数据：")
		var QQ int
		fmt.Scanf("%d",&QQ)  //查询QQ

		startTime := time.Now()
		Qpass,err :=allstrs[QQ]
		if err {
			fmt.Println(QQ,Qpass,"存在")
		}else {
			fmt.Println(QQ,Qpass,"不存在")
		}
		fmt.Println("搜索耗时：",time.Since(startTime))
	}

}
