package main

import (
	"./linking"
	"fmt"
)
func main()  {
	list := linking.NewSingleLinkList()
	node1 := linking.NewSingleLinkNode(1)
	node2 := linking.NewSingleLinkNode(2)
	node3 := linking.NewSingleLinkNode(3)
	list.InsertNodeFront(node1)
	fmt.Println(list)
	list.InsertNodeFront(node2)
	fmt.Println(list)
	list.InsertNodeFront(node3)
	fmt.Println(list)
}
