package main

import (
	"./Queue"
	"fmt"
)
func main()  {
	h := Queue.NewMaxPriQueue()
	h.Insert(*Queue.NewPriItem(101,11))
	h.Insert(*Queue.NewPriItem(102,12))
	h.Insert(*Queue.NewPriItem(103,16))
	h.Insert(*Queue.NewPriItem(104,13))
	h.Insert(*Queue.NewPriItem(105,18))
	fmt.Println(h.Extract())
	fmt.Println(h.Extract())
	fmt.Println(h.Extract())
}
