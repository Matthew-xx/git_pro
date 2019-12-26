package BLC

func (cli *CLI) TestMethod()  {
	blockchain := BlockchainObject()
	defer blockchain.DB.Close()

	utxoSet := &UTXOSet{blockchain}
	utxoSet.ResetUTXOSet()

	//utxoMap := blockchain.FindUTXOMap()
	//fmt.Println(utxoMap)
}