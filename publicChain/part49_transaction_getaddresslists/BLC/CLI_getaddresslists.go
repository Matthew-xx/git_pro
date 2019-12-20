package BLC

import "fmt"

//输出创建的所有钱包地址
func (cli *CLI) addressLists() []string {
	wallets,_ := NewWallets()

	for address,_ := range wallets.WalletsMap{
		fmt.Println(address)

	}
	return nil
}
