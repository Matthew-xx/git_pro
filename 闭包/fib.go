package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

//斐波那契数列
func fibonacci() func() int {
	a ,b := 0,1
	return func() int {
		a,b = b,a+b
		return a
	}
}

type intGen func() int //intGen是一个函数类型，是一个类型就可以实现接口

func (g intGen) Read(p []byte) (n int, err error){
	//函数也能作为接收者，只是一种较特殊的参数
	next := g()  //取得下一个元素，read要把这个元素写进p里面，返回p的字节以及错误信息
	if next > 10000{
		return 0,io.EOF
	} //设置上限，避免无限循环
	s := fmt.Sprintf("%d\n",next) //转换成字符串，把s写进[]bytes
	return strings.NewReader(s).Read(p)  //用NewReader代理reader（即借用已经实现reader接口的）
}

func printFileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main()  {
	f := fibonacci()
	/*  //跟读文件很像，可以将其包装成接口 reader
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	 */
	printFileContents(f)
}