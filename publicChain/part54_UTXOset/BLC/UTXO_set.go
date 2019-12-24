package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

//功能：遍历整个数据库读取所有未花费UTXO，然后将所有的UTXO存储到数据库
//reset重置
//遍历数据库时，返回字典
type UTXOSet struct {
	Blockchain *Blockchain
}

const utxoTableName  = "utxotable"
func (utxoSet *UTXOSet) ResetUTXOSet()  {
	err := utxoSet.Blockchain.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(utxoTableName))

		if b!= nil {
			tx.DeleteBucket([]byte(utxoTableName))  //删除表
			b,_ := tx.CreateBucket([]byte(utxoTableName))

			if b != nil {
				//txOutputsMap := utxoSet.Blockchain.FindUTXOMap()  //未消费输出
			}
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

