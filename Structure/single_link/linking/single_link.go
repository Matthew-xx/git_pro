package linking

import "fmt"

type SingleLink interface {
	//增删改查
	GetFirstNode() *SingleLinkNode  //抓取头部节点
	InsertNode(node *SingleLinkNode)  //头部插入
	InsertBack(node *SingleLinkNode)  //尾部插入
	GetNodeAtIndex(index int) *SingleLinkNode //某些情况下替换原来的值
	DeleteNode(dest *SingleLinkNode)  //删除一个节点
	Deleteatindex(index int )  //删除指定位置的节点
	String() string
 }


type SingleLinkList struct {
	head *SingleLinkNode  //链表头部指针
	length  int   //链表长度
}

func NewSingleLinkList() *SingleLinkList {
	head := NewSingleLinkNode(nil)  //空节点
	return &SingleLinkList{head,0}
}

//返回下一个数据节点
func (list *SingleLinkList) GetFirstNode() *SingleLinkNode {
	return list.head.pNext
}

func (list *SingleLinkList) InsertNodeFront(node *SingleLinkNode)  {
	if list.head == nil {
		list.head.pNext = node
		node.pNext = nil
		list.length++
	}else {
		bak := list.head  //备份
		node.pNext = bak.pNext
		bak.pNext = node
		list.length++   //插入节点，数据追加
	}

}

func (list *SingleLinkList) InsertNodeBack(node *SingleLinkNode)  {
	if list.head == nil {

	}else {

	}
}

func (list *SingleLinkList) String() string {
	var listString string
	p := list.head
	for p.pNext != nil {
		listString += fmt.Sprintf("%v -->",p.pNext.value)
		p = p.pNext  //循环
	}
	listString += fmt.Sprintf("nil")
	return listString  //打印链表字符串
}

