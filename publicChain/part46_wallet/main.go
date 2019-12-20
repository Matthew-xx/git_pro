package main

import (
	"../part46_wallet/BLC"
	"fmt"
)


func main()  {
	wallet := BLC.NewWallet()
	fmt.Printf("%s\n",wallet.GetAddress())

	ISV := wallet.ValidateAddress(string(wallet.GetAddress()))
	fmt.Printf("%v",ISV)
}
