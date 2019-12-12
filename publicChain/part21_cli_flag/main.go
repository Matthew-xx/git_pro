package main

import (
	"flag"
	"fmt"
)

//可以使用命令行+ 参数 运行
//如输入： ./bc printChain   (bc是编译可执行文件，输出printChain函数运行的结果

func main()  {
	flagPrintChainCmd := flag.String("printchain","","输出所有区块：")
	flagInt := flag.Int("INT",88,"输出整数：")
	flagBool := flag.Bool("BOOL",false,"输出布尔值：")

	flag.Parse()

	//  ./main -printchain "gjgjkgui123"
	fmt.Printf("%s\n",*flagPrintChainCmd)

	//  ./main
	fmt.Printf("%d\n",*flagInt)
	fmt.Printf("%v\n",*flagBool)

	//  ./main -printchain "gjgjkgui123" -INT 666 -BOOL

}
