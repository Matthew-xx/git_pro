package main

import (
	"../part47_wallets/BLC"
	"fmt"
)


func main()  {

	wallets := BLC.NewWallets()

	wallets.CreateNewWallet()
	fmt.Println(wallets.Wallets)
	//map[1fGMGnR1UVM4pRrU5woifuJg5mffyBa3F:0xc00009c100]
	// map[key : 钱包对象]，map存储了地址及其对应的对象
}
