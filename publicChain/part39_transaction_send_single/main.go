package main

import (
	".//BLC"
)

//创建创世区块并存到数据库

func main()  {
	//创世区块
	//blockchain := BLC.CreateBlockchainWithGenesisBlock()
	//defer blockchain.DB.Close()

	cli := BLC.CLI{}
	cli.Run()
	// createblockchain -address "mark"  创世块
	// send -from "[\"mark\"]" -to "[\"matthew\"]" -amount "[\"4\"]"
}


