package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

const dbName  = "blockchain.db"
const blockTableName  = "blocks"

//创建区块链结构体
type Blockchain struct {
	Tip []byte    //存储最新区块的hash值
	DB  *bolt.DB    //数据库
	//Blocks []*Block    //存储有序的区块
}

//增加区块到区块链里面
func (blc *Blockchain) AddBlockToBlockchain(data string)  {

	err := blc.DB.Update(func(tx *bolt.Tx) error {
		//获取表
		b := tx.Bucket([]byte(blockTableName))

		if b != nil {
			//获取最新区块并反序列化
			blockBytes := b.Get(blc.Tip)
			block := DeserializeBlock(blockBytes)

			//创建新区块(高度加1
			newBlock := NewBlock(data, block.Height+1, block.Hash)

			//将区块序列化并存储到数据库
			err := b.Put(newBlock.Hash,newBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			//更新数据库里面的“l“对应的hash
			err = b.Put([]byte("l"),newBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			//更新blockchain中的tip
			blc.Tip = newBlock.Hash
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

//创建带有创世区块的区块链
func CreateBlockchainWithGenesisBlock() *Blockchain {

	db,err := bolt.Open(dbName,0600,nil)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	var blockHash []byte

	err = db.Update(func(tx *bolt.Tx) error {
		//获取表（创建
		b := tx.Bucket([]byte(blockTableName))
		if b == nil {
			b ,err = tx.CreateBucket([]byte(blockTableName))
			if err != nil {
				log.Panic(err)
			}
		}

		if b != nil {
			genesisBlock := CreateGenesisBlock("Genesis data")
			//将创世区块存储到库里
			err := b.Put(genesisBlock.Hash,genesisBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}

			err = b.Put([]byte("l"),genesisBlock.Hash) // 存储最新区块的hash
			if err != nil {
				log.Panic(err)
			}

			blockHash = genesisBlock.Hash
		}
		return nil
	})

	//返回区块链对象
	return &Blockchain{blockHash,db}
}



