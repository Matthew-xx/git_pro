package main

import (
	"fmt"
	"os"
)

//可以使用命令行+ 参数 运行
//如输入： ./bc printChain   (bc是编译可执行文件，输出printChain函数运行的结果

func main()  {
	args := os.Args
	fmt.Printf("%v\n",args)
	fmt.Printf("%v\n",args[1]) //输出数组内第2个参数
	//输出一个数组（内容是所有参数的拼接）

}
