package BLC

import (
	"encoding/hex"
	"github.com/boltdb/bolt"
	"log"
)

//功能：遍历整个数据库读取所有未花费UTXO，然后将所有的UTXO存储到数据库
//reset重置
//遍历数据库时，返回字典
//存储到数据库
type UTXOSet struct {
	Blockchain *Blockchain
}

const utxoTableName  = "utxotable"

func (utxoSet *UTXOSet) ResetUTXOSet()  {
	err := utxoSet.Blockchain.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(utxoTableName))

		if b!= nil {
			err := tx.DeleteBucket([]byte(utxoTableName))  //删除表
			if err != nil {
				log.Panic(err)
			}
		}

		b,_ = tx.CreateBucket([]byte(utxoTableName)) //不管有没有库都创建
		if b != nil {
			txOutputsMap := utxoSet.Blockchain.FindUTXOMap()  //获取所有未消费输出

			for keyHash,outs := range txOutputsMap{
				txHash,_ := hex.DecodeString(keyHash)
				b.Put(txHash,outs.Serialize())
			}
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

