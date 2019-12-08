package main

import (
	"../part2_Basic_Prototype/BLC"
	"fmt"
)

func main()  {
	genesisBlock := BLC.CreateGenesisBlock("genesis block")
	fmt.Println(genesisBlock)
}


