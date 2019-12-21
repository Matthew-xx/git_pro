package BLC

func (cli *CLI) createGenesis(address string) {
	//创建coinbase的交易
	blockchain := CreateBlockchainWithGenesisBlock(address)
	defer blockchain.DB.Close()
}

