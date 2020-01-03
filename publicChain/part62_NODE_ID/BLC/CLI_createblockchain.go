package BLC


// 创建创世区块
func (cli *CLI) createGenesisBlockchain(address string,nodeID string)  {

	blockchain := CreateBlockchainWithGenesisBlock(address,nodeID)
	defer blockchain.DB.Close()

	//在创建创世区块的时候就将UTXO存储到未消费表中
	utxoSet := &UTXOSet{blockchain}
	utxoSet.ResetUTXOSet()
}