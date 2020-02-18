package main

import "fmt"

func main1()  {
	node1 := new(Node)
	node2 := new(Node)
	node3 := new(Node)
	node1.data = 1
	node1.pNext = node2
	node2.data = 2
	node2.pNext = node3
	node3.data = 3
	fmt.Println(node1.pNext.pNext.data)
}

func main2()  {
	mystack := NewStack()
	for i:=0;i<100;i++ {
		mystack.Push(i)
	}
	for data:=mystack.Pop();data!=nil;data=mystack.Pop() {
		fmt.Println(data)
	}
}
func main()  {
	myq := NewLinkQueue()
	for i:=0;i<100;i++ {
		myq.Enqueue(i)
	}
	for data:=myq.Dnqueue();data!=nil;data=myq.Dnqueue() {
		fmt.Println(data)
	}
}
