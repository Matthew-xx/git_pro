package BLC

import (
	"bytes"
	"fmt"
)

type TxOutput struct {
	Value int64
	hashPubKey []byte  //公钥（一次256一次160hash后的HashPubKey
}

//上锁(设置160hash值hashPubKey
//func (txOutput *TxOutput)  Lock(address string)  {
//
//	publicKeyHash := Base58Decode([]byte(address))
//
//	txOutput.hashPubKey = publicKeyHash[1:len(publicKeyHash) - 4]
//}
/*
func (txOutput *TxOutput) Lock(address string)  {
	publickeyHash := Base58Decode([]byte(address))  //得到25字节数组
	txOutput.hashPubKey = publickeyHash[1:len(publickeyHash)-4]  //除去版本及末尾4位
}*/

func NewTxOutput(value int64,address string) *TxOutput {


	publicKeyHash := Base58Decode([]byte(address))
	hashPubKey := publicKeyHash[1:len(publicKeyHash) - 4]
	//fmt.Println(hashPubKey)

	txOutput := &TxOutput{value,hashPubKey}
	fmt.Println(txOutput)

	//设置hashPubKey
	//txOutput.Lock(address)
	return txOutput
}

//解锁(判断转账的地址与txouput是否一致
func (txOutput *TxOutput) UnlockScriptPubKeyWithAddress(address string) bool {
	publickeyHash := Base58Decode([]byte(address))
	Hash160 := publickeyHash[1:len(publickeyHash)-4]

	return bytes.Compare(txOutput.hashPubKey,Hash160) == 0
}



