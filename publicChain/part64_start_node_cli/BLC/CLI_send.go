package BLC

import (
	"fmt"
	"os"
)

// 转账
func (cli *CLI) send(from []string,to []string,amount []string,nodeID string,mineNow bool)  {


	if DBExists() == false {
		fmt.Println("数据不存在.......")
		os.Exit(1)
	}

	blockchain := BlockchainObject(nodeID)
	defer blockchain.DB.Close()

	if mineNow {
		blockchain.MineNewBlock(from,to,amount,nodeID)

		utxoSet := &UTXOSet{blockchain}

		//转账成功以后，需要更新一下
		utxoSet.Update()
	} else {
		// 把交易发送到矿工节点去进行验证
		fmt.Println("由矿工节点处理......")
	}

}