package main

import (
	"../part3_Basic_Prototype/BLC"
	"fmt"
)

func main()  {
	genesisBlockchain := BLC.CreateBlockchainWithGenesisBlock()
	fmt.Println(genesisBlockchain)  //区块链地址
	fmt.Println(genesisBlockchain.Block)
	fmt.Println(genesisBlockchain.Block[0])  //打印第一个区块（创世区块
}


