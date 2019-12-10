package main

import (
	"../part7_proof_of_work/BLC"
	"fmt"
)

func main()  {
	//创世区块
	blockchain := BLC.CreateBlockchainWithGenesisBlock()

	//fmt.Println(blockchain.Blocks[0])  //打印第一个区块（创世区块
	//新区块
	blockchain.AddBlockToBlockchain("send 100RMB to mark",
		blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,
		blockchain.Blocks[len(blockchain.Blocks)-1].Hash) //第二、三个参数为当前区块的高度和hash值
	blockchain.AddBlockToBlockchain("send 50RMB to matthew",
		blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,
		blockchain.Blocks[len(blockchain.Blocks)-1].Hash) //第二、三个参数为当前区块的高度和hash值

	fmt.Println(blockchain)  //区块链地址
	//fmt.Println(blockchain.Blocks)
}


