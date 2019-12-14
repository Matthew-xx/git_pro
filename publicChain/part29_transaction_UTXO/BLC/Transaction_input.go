package BLC

type TxInPut struct {
	Txid []byte  //交易的ID（hash值）
	Vout int    //存储在TXoutput在Vout里面的索引
	ScriptSig  string //用户名
}
