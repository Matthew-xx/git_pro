package BLC

import "bytes"

type TxInPut struct {
	Txid []byte  //交易的ID（hash值）
	Vout int    //存储在TXoutput在Vout里面的索引
	Signature []byte  //数字签名
	PublicKey   []byte   //公钥（钱包里的原始公钥
}

//判断当前的消费属于哪个用户
func (txInput *TxInPut) UnlockPubKey(hashPubKey []byte) bool {
	//PubKey 是out里的公钥
	publicKey := HashPubKey(txInput.PublicKey)

	return bytes.Compare(publicKey,hashPubKey)==0
}
