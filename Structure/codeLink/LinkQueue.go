package main

type QueueLink struct {
	rear *Node   //指向头节点
	front *Node  //指向尾节点
}

type LinkQueue interface {
	Length() int
	Enqueue(value interface{})  //入队
	Dnqueue() interface{}  //出队
}

func NewLinkQueue() *QueueLink {
	return &QueueLink{}
}

func (qlk *QueueLink) Length() int{
	pnext := qlk.front
	length := 0
	for pnext.pNext != nil {
		pnext = pnext.pNext //节点的循环跳跃
		length++
	}
	return length
}

//头部插入
func (qlk *QueueLink) Enqueue(value interface{}){
	newnode := &Node{value,nil}
	if qlk.front == nil {
		qlk.front = newnode  //插入一个节点
		qlk.rear = newnode
	}else{
		qlk.rear.pNext = newnode
		qlk.rear = qlk.rear.pNext  //
	}
}

func (qlk *QueueLink) Dnqueue() interface{} {
	if qlk.front == nil {
		return nil
	}
	newnode := qlk.front  //记录头部位置
	if qlk.front == qlk.rear { //只有一个的情况
		qlk.front = nil
		qlk.rear = nil
	}else {
		qlk.front = qlk.front.pNext  //删除一个
	}
	return newnode.data
}

