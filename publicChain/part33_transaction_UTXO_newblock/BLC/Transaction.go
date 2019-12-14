package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

//UTXO（未花费输出
type Transaction struct {
	//交易hash
	TxHash []byte
	//输入
	Vins []*TxInPut
	//输出
	Vouts []*TxOutput

}

//交易（transaction）的创建分两种情况
//1. 创世区块创建时的transaction
//2. 转账时产生的transaction

//1. 创世区块第一笔交易
func NewCoinbaseTransaction(address string) *Transaction {
	//创世区块输入(代表消费)
	txInput := &TxInPut{[]byte{},-1,"Genesis Data"}
	//代表未消费
	txOutput := &TxOutput{10,address}
	//交易
	txCoinbase := &Transaction{[]byte{},[]*TxInPut{txInput},[]*TxOutput{txOutput}}
	//设置hash值
	txCoinbase.HashTxSerialize()

	return txCoinbase
}




//序列化后再Hash
func (tx *Transaction) HashTxSerialize()  {

	var result bytes.Buffer  //定义缓冲
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	hash := sha256.Sum256(result.Bytes())
	tx.TxHash = hash[:]
}


