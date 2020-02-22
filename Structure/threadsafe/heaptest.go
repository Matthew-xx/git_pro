package main

import (
	"./Queue"
	"fmt"
)




func main()  {
	h := Queue.NewMin()
	h.Insert(Queue.Int(8))
	h.Insert(Queue.Int(9))
	h.Insert(Queue.Int(7))
	h.Insert(Queue.Int(10))
	h.Insert(Queue.Int(6))
	h.Insert(Queue.Int(2))
	h.Insert(Queue.Int(3))
	fmt.Println(h.Extract().(Queue.Int))
	fmt.Println(h.Extract().(Queue.Int))
}
