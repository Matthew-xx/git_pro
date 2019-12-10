package main

import (
	"../part14_block_bolt/BLC"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)
//将block序列化存储到数据库

func main()  {
	/*
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
	 */
	block := BLC.NewBlock("Test",1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
	fmt.Printf("%d\n",block.Nonce)
	fmt.Printf("%x\n",block.Hash)

	//先验证（上面):hash有效，再将其当做参数传入来判断看是否真
	//proofOfWork := BLC.NewProofOfWork(block)
	//fmt.Printf("%v",proofOfWork.IsValid())

	/*
	bytes := block.Serialize()  //block的行为，调用函数
	fmt.Println(bytes)

	block = BLC.DeserializeBlock(bytes)

	fmt.Printf("%d\n",block.Nonce)
	fmt.Printf("%x\n",block.Hash)
	 */
	//创建或者打开数据库
	db, err := bolt.Open("my.db",0600,nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//创建或更新表
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))

		if b == nil {
			b,err := tx.CreateBucket([]byte("blocks"))
			if err != nil {
				log.Panic("blocks table create fail",err)
			}
			b = b
		}
		//将block序列化并存储
		err = b.Put([]byte("l"),block.Serialize())
		if err!= nil {
			log.Panic(err)
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	//查看表
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))
		if b != nil {
			blockData := b.Get([]byte("l"))
			fmt.Printf("%s\n",blockData)
			//打印序列化和反序列化
			block := BLC.DeserializeBlock(blockData)
			fmt.Println(block)
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}


