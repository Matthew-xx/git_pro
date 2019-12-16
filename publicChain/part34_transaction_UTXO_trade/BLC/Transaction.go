package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
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

func NewSimpleTransaction(from string,to string,amount int) *Transaction {

	var txInputs []*TxInPut
	var txOutputs []*TxOutput

	//输入(代表消费)
	bytes ,_ := hex.DecodeString("887ec7801626766c488f3b09f16ae34571219150be16d354db36a17c2ce45dd4")
	txInput := &TxInPut{bytes,0,from}
	txInputs = append(txInputs,txInput)

	//输出代表未消费
	txOutput := &TxOutput{int64(amount),to}
	txOutputs = append(txOutputs,txOutput)

	//找零
	txOutput = &TxOutput{10-int64(amount),from}
	txOutputs = append(txOutputs,txOutput)

	//交易
	tx := &Transaction{[]byte{},txInputs,txOutputs}
	//设置hash值
	tx.HashTxSerialize()

	return tx
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


