package BLC

func (cli *CLI) TestMethod(nodeID string)  {
	blockchain := BlockchainObject(nodeID)
	defer blockchain.DB.Close()

	utxoSet := &UTXOSet{blockchain}
	utxoSet.ResetUTXOSet()

	//utxoMap := blockchain.FindUTXOMap()
	//fmt.Println(utxoMap)
}