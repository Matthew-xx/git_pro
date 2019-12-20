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
	fmt.Println("\tcreatewallet  -- 创建钱包")
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

	createWalletCmd := flag.NewFlagSet("createwallet",flag.ExitOnError)
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
	case "createwallet":
		err := createWalletCmd.Parse(os.Args[2:])  //解析addBlock后面的内容
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

		//在发送（转账)的时候判断地址是否有效
		from := JsonToArray(*flagFrom)
		to := JsonToArray(*flagTo)

		for index, fromAddress := range from{
			if IsValidateAddress([]byte(fromAddress)) == false || IsValidateAddress([]byte(to[index])) == false{
				//from 和 to 的地址都应该有效
				fmt.Println("地址无效...")
				printUsage()
				os.Exit(1)
			}

		}
		amount := JsonToArray(*flagAmount)
		cli.Send(from,to,amount)
	}

	if printChainCmd.Parsed() {
		//fmt.Println("输出所有区块的数据")
		cli.printchain()
	}

	if createWalletCmd.Parsed() {
		//创建钱包
		cli.createWallet()
	}

	if createGenesisCmd.Parsed() {
		if IsValidateAddress([]byte(*flagcreateGenesisAddress)) == false{
			fmt.Println("地址无效...")
			printUsage()
			os.Exit(1)
		}
		cli.createGenesis(*flagcreateGenesisAddress)
	}

	if getbalanceCmd.Parsed() {
		if IsValidateAddress([]byte(*getbalanceWithAddress)) == false {
			fmt.Println("地址无效...")
			printUsage()
			os.Exit(1)
		}
		cli.getBalance(*getbalanceWithAddress)
	}
}

