package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

//迭代器
type BlockchainIterator struct {
	CurrentHash  []byte
	DB  *bolt.DB
}

func (blockchainIterator *BlockchainIterator) Next() *Block {

	var block *Block
	err := blockchainIterator.DB.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			currentBlockBytes := b.Get(blockchainIterator.CurrentHash)
			//获取当前迭代器里的currenthash所对应的区块
			block = DeserializeBlock(currentBlockBytes)
			//更新迭代器里面的currenthash
			blockchainIterator.CurrentHash = block.PrevBlockHash
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return block
}