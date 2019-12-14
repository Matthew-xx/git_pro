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
	fmt.Println("\taddblock -data DATA -- 交易数据")
	fmt.Println("\tprintchain --输出区块信息")
}

func isValidArgs()  {
	if len(os.Args) <2 {
		printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addblock(txs []*Transaction)  {
	if DbExists() == false{
		fmt.Println("数据库不存在")
		os.Exit(1)
	}

	blockchain := GetBlockChainObject()
	defer blockchain.DB.Close()

	blockchain.AddBlockToBlockchain([]*Transaction{})
}

func (cli *CLI) printchain()  {
	if DbExists() == false{
		fmt.Println("数据库不存在")
		os.Exit(1)
	}

	blockchain := GetBlockChainObject()
	defer blockchain.DB.Close()

	blockchain.PrintChain()
}

func (cli *CLI) createGenesis(address string) {
	//创建coinbase的交易
	CreateBlockchainWithGenesisBlock(address)
}

func (cli *CLI) Run() {
	isValidArgs()

	addBlockCmd := flag.NewFlagSet("addblock",flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain",flag.ExitOnError)
	createGenesisCmd := flag.NewFlagSet("address",flag.ExitOnError)

	flagAddBlockData := addBlockCmd.String("data","www.bgsgg.com","交易数据")
	flagcreateGenesisAddress := createGenesisCmd.String("address","","创世区块地址")
	//传入地址后hash化（coinbase.transaction）再传入到第一个区块

	//看第一个参数是什么功能
	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])  //解析addBlock后面的内容
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
		cli.addblock([]*Transaction{})
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
}

