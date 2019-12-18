package BLC

//返回的未消费输出，（他本身及所在的位置
type UTXO struct {
	TxHash []byte  //代表当前未消费输出所在的transaction的hash
	Index int      //transaction中的Vout中有多个输出，索引
	Output *TxOutput  //他自己本身
}
