package main

import (
	"./hash_Arr"
	"fmt"
)

func main1()  {
	fmt.Println(hash_Arr.MySHA("abcd",100))
	fmt.Println(hash_Arr.MySHA("abcd1",100))
	fmt.Println(hash_Arr.MySHA256("abcd",100))
	fmt.Println(hash_Arr.MySHA256("abcd",100))
}

func main()  {
	table,_ := hash_Arr.NewHashTable(100,hash_Arr.MySHA)
	table.Insert("abcd")
	table.Insert("abcd1")
	pos := table.Find("abcd")
	fmt.Println(table.GetValue(pos))
}
