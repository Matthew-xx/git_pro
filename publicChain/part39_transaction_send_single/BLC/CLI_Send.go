package BLC

import (
	"fmt"
	"os"
)

//转账功能
func (cli *CLI)Send(from []string,to []string,amount []string)  {
	if DbExists() == false{
		fmt.Println("数据库不存在")
		os.Exit(1)
	}

	//先获取blockchain对象，然后进行操作MineNewBlock，最后关闭数据库
	blockchain := GetBlockChainObject()
	defer blockchain.DB.Close()

	blockchain.MineNewBlock(from,to,amount)
}
