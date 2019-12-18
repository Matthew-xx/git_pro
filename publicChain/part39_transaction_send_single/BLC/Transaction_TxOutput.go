package BLC


type TxOutput struct {
	Value int64
	ScriptPubKey string  //用户名
}


//解锁
func (txOutput *TxOutput) UnlockScriptPubKeyWithAddress(address string) bool {

	return txOutput.ScriptPubKey == address
}
