package BLC

import "fmt"

// 先用它去查询余额
func (cli *CLI) getBalance(address string,nodeID string)  {

	//fmt.Println("地址：" + address)

	blockchain := BlockchainObject(nodeID)
	defer blockchain.DB.Close()

	utxoSet := &UTXOSet{blockchain}
	amount := utxoSet.GetBalance(address)  //不需遍历整个数据库

	//amount := blockchain.GetBalance(address)  //遍历整个数据库查找

	fmt.Printf("%s 一共有%d 个Token\n",address,amount)


}
