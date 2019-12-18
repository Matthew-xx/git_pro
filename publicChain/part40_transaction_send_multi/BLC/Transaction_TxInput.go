package BLC

type TxInPut struct {
	Txid []byte  //交易的ID（hash值）
	Vout int    //存储在TXoutput在Vout里面的索引
	ScriptSig  string //用户名
}

//判断当前的消费属于哪个用户
func (txInput *TxInPut) UnlockScriptSigWithAddress(address string) bool {

	return txInput.ScriptSig == address
}
