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

func (utxoSet *UTXOSet) findUTXOForAddress(address string) []*UTXO {
	var utxos []*UTXO

	utxoSet.Blockchain.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(utxoTableName))

		c := b.Cursor()
		//k 是transaction的hash，v 对应outputs对象
		for k,v := c.First(); k!= nil; k,v = c.Next() {
			//fmt.Printf("key=%x , value=%x\n",k,v)
			txOutputs :=DeserializeTXOutputs(v)

			for _,utxo := range txOutputs.UTXOS{
				//判断地址是否匹配（解锁
				if utxo.Output.UnLockScriptPubKeyWithAddress(address) {
					utxos = append(utxos,utxo)
				}
			}
		}
		return nil
	})
	return utxos
}

func (utxoSet *UTXOSet) GetBalance(address string) int64 {
	//找到所有未花费输出
	UTXOS := utxoSet.findUTXOForAddress(address)

	var amount int64
	for _,utxo := range UTXOS{
		amount += utxo.Output.Value
	}

	return amount
}

