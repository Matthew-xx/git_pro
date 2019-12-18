package BLC

import (
	"log"
	"os"
	"fmt"
	"flag"
)

type CLI struct {

}

func printUsage()  {
	fmt.Println("Usage:")
	fmt.Println("\tcreategenesis -address -- 创建创世区块")
	fmt.Println("\tsend -from FROM -to TO -amount AMOUNT -- 转账明细")
	fmt.Println("\tprintchain --输出区块信息")
	fmt.Println("\tgetbalance -address --输出地址未消费输出")
}

func isValidArgs()  {
	if len(os.Args) <2 {
		printUsage()
		os.Exit(1)
	}
}


func (cli *CLI) Run() {
	isValidArgs()

	sendBlockCmd := flag.NewFlagSet("send",flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain",flag.ExitOnError)
	createGenesisCmd := flag.NewFlagSet("address",flag.ExitOnError)
	getbalanceCmd := flag.NewFlagSet("getbalance",flag.ExitOnError)

	//flagAddBlockData := sendBlockCmd.String("data","www.bgsgg.com","交易数据")
	flagFrom := sendBlockCmd.String("from","","转账地址")
	flagTo := sendBlockCmd.String("to","","转账目的地址")
	flagAmount := sendBlockCmd.String("amount","","转账金额")

	flagcreateGenesisAddress := createGenesisCmd.String("address","","创世区块地址")
	//传入地址后hash化（coinbase.transaction）再传入到第一个区块

	getbalanceWithAddress := getbalanceCmd.String("address","","查询某账号余额")

	//看第一个参数是什么功能
	switch os.Args[1] {
	case "send":
		err := sendBlockCmd.Parse(os.Args[2:])  //解析addBlock后面的内容
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])  //解析printchain后面的内容
		if err != nil {
			log.Panic(err)
		}
	case "creategenesis":
		err := createGenesisCmd.Parse(os.Args[2:])  //解析addBlock后面的内容
		if err != nil {
			log.Panic(err)
		}
	case "getbalance":
		err := getbalanceCmd.Parse(os.Args[2:])  //解析addBlock后面的内容
		if err != nil {
			log.Panic(err)
		}
	default:
		printUsage()
		os.Exit(1)
	}

	//解析成功
	if sendBlockCmd.Parsed() {
		if *flagFrom == "" || *flagTo == "" || *flagAmount == "" {
			printUsage()
			os.Exit(1)
		}

		from := JsonToArray(*flagFrom)
		to := JsonToArray(*flagTo)
		amount := JsonToArray(*flagAmount)
		cli.Send(from,to,amount)
	}

	if printChainCmd.Parsed() {
		//fmt.Println("输出所有区块的数据")
		cli.printchain()
	}

	if createGenesisCmd.Parsed() {
		if *flagcreateGenesisAddress == "" {
			fmt.Println("地址不能为空...")
			printUsage()
			os.Exit(1)
		}
		cli.createGenesis(*flagcreateGenesisAddress)
	}

	if getbalanceCmd.Parsed() {
		if *getbalanceWithAddress == "" {
			fmt.Println("地址不能为空...")
			printUsage()
			os.Exit(1)
		}
		cli.getBalance(*getbalanceWithAddress)
	}
}

