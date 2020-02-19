package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"strings"
)

func main()  {
	fi,err := os.Open("F:\\Software\\go_path\\my_pro\\Structure\\data\\data.txt")
	if err != nil {
		errors.New("文件读取失败")
		return
	}
	defer fi.Close()

	path := "F:\\Software\\go_path\\my_pro\\Structure\\data\\datasort.txt"  //保存文件的路径
	savefile,_ := os.Create(path)
	defer savefile.Close()

	savebuff := bufio.NewWriter(savefile)  //用于写入的对象

	br := bufio.NewReader(fi)
	for {  //循环，逐行读取数据
		line,_,err := br.ReadLine() //读取一行
		if err == io.EOF {
			break  //跳出循环
		}
		//fmt.Println(string(line))
		linestr := string(line)  //转换为字符串
		mystrs := strings.Split(linestr,"  #  ")  //字符串切割
		fmt.Println(mystrs[1])
		fmt.Fprintln(savebuff,mystrs[1])  //打印到文件
	}
	savebuff.Flush()
}

