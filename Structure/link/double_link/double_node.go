package double_link

//双向链表节点
type DoubleLinkNode struct {
	value interface{}
	prev *DoubleLinkNode  //上一个节点
	next *DoubleLinkNode  //下一个节点
}

//新建一个节点
func NewDoubleLinkNode(value interface{}) *DoubleLinkNode {
	return &DoubleLinkNode{value,nil,nil}
}
//返回数据
func (node *DoubleLinkNode) Value() interface{} {
	return node.value
}
//返回上一个节点数据
func (node *DoubleLinkNode) Prev() *DoubleLinkNode {
	return node.prev
}
//返回下一个节点数据
func (node *DoubleLinkNode) Pnext() *DoubleLinkNode {
	return node.next
}



