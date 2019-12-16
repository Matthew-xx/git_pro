package main

import (
	"../part36_transaction_UTXO_balance/BLC"
)

//创建创世区块并存到数据库

func main()  {
	//创世区块
	//blockchain := BLC.CreateBlockchainWithGenesisBlock()
	//defer blockchain.DB.Close()

	cli := BLC.CLI{}
	cli.Run()

	// send -from "[\"mark\"]" -to "[\"matthew\"]" -amount "[\"4\"]"
}


