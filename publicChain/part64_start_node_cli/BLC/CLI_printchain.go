package BLC

import (
	"fmt"
	"os"
)

func (cli *CLI) printchain(nodeID string)  {

	if DBExists() == false {
		fmt.Println("数据不存在.......")
		os.Exit(1)
	}

	blockchain := BlockchainObject(nodeID)

	defer blockchain.DB.Close()

	blockchain.Printchain()

}