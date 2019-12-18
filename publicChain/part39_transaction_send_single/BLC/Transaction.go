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

//判断当前交易是否创世区块交易
func (tx *Transaction) IsCoinbaseTransaction() bool {

	//判断input里面的内容
	return len(tx.Vins[0].Txid) == 0 && tx.Vins[0].Vout == -1
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

func NewSimpleTransaction(from string,to string,amount int,blockchain *Blockchain) *Transaction {

	//函数实现返回from用户所有的未花费输出对应的transaction

	//unSpentTx := blockchain.UnUTXOs(from)
	//fmt.Println(unSpentTx)


	//函数实现，返回未消费余额及其hash（字典形式

	money,spentableUTXOdic := blockchain.FindSpentableUTXOs(from,amount)

	var txInputs []*TxInPut
	var txOutputs []*TxOutput

	//输入(代表消费)
	/*
	bytes ,_ := hex.DecodeString("887ec7801626766c488f3b09f16ae34571219150be16d354db36a17c2ce45dd4") //先转成字节数组
	txInput := &TxInPut{bytes,0,from}  // 0 是索引（如上一个区块有两个输出，则0表示消费第一个输出
	txInputs = append(txInputs,txInput)
	 */
	//如遍历{hash1:[0],hash2[2,3]}，则第一个for循环遍历2次（2个键值对hash1,hash2),
	for txHash,indexArray := range  spentableUTXOdic{
		txHashBytes,_ := hex.DecodeString(txHash)
		for _,index := range indexArray{
			txInput := &TxInPut{txHashBytes,index,from}
			txInputs = append(txInputs,txInput)
		}
	}

	//输出,代表未消费(转账
	txOutput := &TxOutput{int64(amount),to}
	txOutputs = append(txOutputs,txOutput)

	//找零
	txOutput = &TxOutput{int64(money)-int64(amount),from}
	txOutputs = append(txOutputs,txOutput)

	//交易
	tx := &Transaction{[]byte{},txInputs,txOutputs}

	tx.HashTxSerialize() //设置hash值

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


