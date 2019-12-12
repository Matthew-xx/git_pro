package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

//可以使用命令行+ 参数 运行
//如输入： ./bc printChain   (bc是编译可执行文件，输出printChain函数运行的结果

func printUsage()  {
	fmt.Println("Usage:")
	fmt.Println("\taddblock -data DATA -- 交易数据")
	fmt.Println("\tprintchain --输出区块信息")
}

func isValidArgs()  {
	if len(os.Args) <2 {
		printUsage()
		os.Exit(1)
	}
}

func main()  {
	isValidArgs()

	addBlockCmd := flag.NewFlagSet("addblock",flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain",flag.ExitOnError)

	flagAddBlockData := addBlockCmd.String("data","www.bgsgg.com","交易数据")

	//看第一个参数是什么功能
	switch os.Args[1] {
	case "addBlock":
		err := addBlockCmd.Parse(os.Args[2:])  //解析addBlock后面的内容
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])  //解析printchain后面的内容
		if err != nil {
			log.Panic(err)
		}
	default:
		printUsage()
		os.Exit(1)
	}

	//解析成功
	if addBlockCmd.Parsed() {
		if *flagAddBlockData == ""{
			printUsage()
			os.Exit(1)
		}
		fmt.Println(*flagAddBlockData)
	}

	if printChainCmd.Parsed() {
		fmt.Println("输出所有区块的数据")
	}
}


//  ./main addblock -data "maxxxxsf"   :输出  maxxxxxsf