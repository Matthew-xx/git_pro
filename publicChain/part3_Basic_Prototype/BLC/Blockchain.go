package BLC

//创建区块链结构体
type Blockchain struct {
	Block []*Block    //存储有序的区块
}

//创建带有创世区块的区块链
func CreateBlockchainWithGenesisBlock() *Blockchain {
	//创建创世区块
	genesisBlock := CreateGenesisBlock("Genesis data")
	//返回区块链对象
	return &Blockchain{[]*Block{genesisBlock}}
}



