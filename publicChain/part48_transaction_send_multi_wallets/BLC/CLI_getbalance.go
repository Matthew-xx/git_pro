package BLC

import (
	"fmt"
)
//查询余额
func (cli *CLI) getBalance(address string)  {
	fmt.Println("地址："+ address)

	blockchain := GetBlockChainObject()
	defer blockchain.DB.Close()

	amount := blockchain.GetBalance(address)

	fmt.Printf("%s 共有 %d 个token\n",address,amount)
}

