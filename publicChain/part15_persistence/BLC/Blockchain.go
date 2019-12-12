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
//func (blc *Blockchain) AddBlockToBlockchain(data string,height int64,preHash []byte)  {
//	//创建新区块
//	newBlock := NewBlock(data,height,preHash)
//	//往链里面添加区块
//	blc.Blocks = append(blc.Blocks,newBlock)
//}

//创建带有创世区块的区块链
func CreateBlockchainWithGenesisBlock() *Blockchain {

	db,err := bolt.Open(dbName,0600,nil)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	var blockHash []byte

	err = db.Update(func(tx *bolt.Tx) error {
		b ,err := tx.CreateBucket([]byte(blockTableName))
		if err != nil {
			log.Panic(err)
		}
		if b == nil {
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

	//创建创世区块

	//返回区块链对象
	return &Blockchain{blockHash,db}
}



