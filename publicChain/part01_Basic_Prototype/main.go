package main

import (
	"../part1_Basic_Prototype/BLC"
	"fmt"
)

func main()  {
	block := BLC.NewBlock("马雄星",1,[]byte{0,0,0,0,0})
	fmt.Println(block)
}


