package main

type Node struct {
	data interface{}
	pNext *Node  //指针指向下一个节点
}

type LinkStack interface {
	IsEmpty() bool
	Push(data interface{})
	Pop()  interface{}
	Length()  int
}

func NewStack() *Node  {
	return &Node{} //返回一个节点指针
}

func (n *Node) IsEmpty() bool{
	if n.pNext == nil {
		return true
	}else {
		return false
	}
}

func (n *Node) Push(data interface{}){
	newnode := &Node{data,nil} //
	newnode.pNext = n.pNext  //插入
	n.pNext = newnode //头部插入
}

//栈的操作方向是一样
func (n *Node) Pop() interface{}{
	if n.IsEmpty() == true {
		return nil
	}  //头部操作
	value := n.pNext.data  //要弹出的数据
	n.pNext = n.pNext.pNext //删除
	return value
}

func (n *Node) Length() int{
	pnext := n
	length := 0
	for pnext.pNext != nil {
		pnext = pnext.pNext //节点的循环跳跃
		length++
	}
	return length
}


