package linking

//单链表，每个元素存储到下一个位置

//链表的节点
type SingleLinkNode struct {
	value  interface{}
	pNext  *SingleLinkNode
}


//构造节点
func NewSingleLinkNode(data interface{}) *SingleLinkNode {
	return &SingleLinkNode{data,nil}
}

//返回节点数据
func (node *SingleLinkNode) Value() interface{} {
	return node.value
}

//返回下一节点地址
func (node *SingleLinkNode) PNext() *SingleLinkNode {
	return node.pNext
}

