package main

//顺序搜索
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type QQ struct {
	Quser   int
	Qpass   string
	Qname   string
}

func main()  {

	const N = 27  //需要开辟的内存
	allstrs := make([]QQ,N,N)  //开辟数组存储数据
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

	fmt.Println("要搜索的数据：")
	var QQ int
	fmt.Scanf("%d",&QQ)  //输入要搜索的数据
	//fmt.Scanln(&QQ)

	startTime := time.Now()
	for j:=0;j<N;j++ {
		if allstrs[j].Quser == QQ { //搜索字符串  “h"
			fmt.Println(j,allstrs[j].Quser,allstrs[j].Qpass,allstrs[j].Qname) //根据QQ查询密码
		}
	}
	fmt.Println("搜索耗时：",time.Since(startTime))
}

