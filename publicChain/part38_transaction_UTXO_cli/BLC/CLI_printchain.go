package BLC

import (
	"fmt"
	"os"
)

//打印数据库中区块
func (cli *CLI) printchain()  {
	if DbExists() == false{
		fmt.Println("数据库不存在")
		os.Exit(1)
	}

	blockchain := GetBlockChainObject()
	defer blockchain.DB.Close()

	blockchain.PrintChain()
}
