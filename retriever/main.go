package main

import (
	"./baidu"
	"fmt"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}
//download是使用者，使用者要Get，因此要定义一个接口interface的接收者retriever，接口里面有Get(不需要加func关键字，接口里面本身就是函数）
//然后baiduretriever实现了Get方法，也即实现了retriever接口
func main() {
	var r Retriever  //r还没定义，接下来实现一个interface
	r = baidu.Retriever{}  //通过值
	//r = &baidu.Retriever{}  //通过指针
	//fmt.Println(download(r))
	inspect(r)

	//接口变量里面有两个东西：实现者的类型和实现者的值
	//获取接口的值类型，方法一：type assertion
	retrieverType := r.(baidu.Retriever)
	fmt.Println(retrieverType.TimeOut)
}
// 获取值类型方法二，
func inspect(r Retriever)  {
	fmt.Println("%T %v",r,r)
	switch v := r.(type) {
	case baidu.Retriever:
		fmt.Println("TimeOut:",v.TimeOut)
	}
}