package main

import (
	"../part49_transaction_getaddresslists/BLC"
)

//创建创世区块并存到数据库

func main()  {
	//创世区块
	//blockchain := BLC.CreateBlockchainWithGenesisBlock()
	//defer blockchain.DB.Close()

	cli := BLC.CLI{}
	cli.Run()

    //createblockchain -address "mark"
	// send -from "[\"mark\"]" -to "[\"matthew\"]" -amount "[\"4\"]"
	//send -from "[\"mark\",\"maxin\"]" -to "[\"maxin\",\"mark\"]" -amount "[\"1\",\"3\"]"
}


