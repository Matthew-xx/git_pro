package main

//硬盘搜索
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main()  {
	startTime := time.Now()
	path := "F:\\Software\\go_path\\my_pro\\Structure\\data\\data.txt"
	file,_ := os.Open(path)
	defer file.Close()

	br := bufio.NewReader(file)
	i := 0  //统计共多少行
	for  {
		line,_,end := br.ReadLine()  //逐行读取
		if end == io.EOF { //文件关闭
			break
		}
		linestr := string(line)
		if strings.Contains(linestr,"星") { //搜索字符串  “h"
			fmt.Println(linestr)                 //是硬盘搜索
			i++
		}
	}
	fmt.Println("搜索耗时：",time.Since(startTime))
	fmt.Printf("共发现 %d 行数据",i)
}
