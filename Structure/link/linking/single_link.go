package linking

//头部插入的数据较尾插快

import (
	"fmt"
	"strings"
)

type SingleLink interface {
	//增删改查
	GetFirstNode() *SingleLinkNode  //抓取头部节点
	InsertNode(node *SingleLinkNode)  //头部插入
	InsertBack(node *SingleLinkNode)  //尾部插入

	//在一个节点之前或之后插入
	InsertNodeValueBack(dest interface{},node *SingleLinkNode)  bool
	InsertNodeValueFront(dest interface{},node *SingleLinkNode)  bool

	GetNodeAtIndex(index int) *SingleLinkNode //某些情况下替换原来的值
	DeleteNode(dest *SingleLinkNode) bool //删除一个节点
	Deleteatindex(index int )  //删除指定位置的节点
	String() string
	FindString(data string)

	ReverseList()
 }


type SingleLinkList struct {
	head *SingleLinkNode  //链表头部指针
	length  int   //链表长度
}

func NewSingleLinkList() *SingleLinkList {
	head := NewSingleLinkNode(nil)  //空节点
	return &SingleLinkList{head,0}
}

//返回第一个数据节点
func (list *SingleLinkList) GetFirstNode() *SingleLinkNode {
	return list.head.pNext
}

//头部插入
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

//尾部插入
func (list *SingleLinkList) InsertNodeBack(node *SingleLinkNode)  {
	if list.head == nil {
		list.head.pNext = node
		node.pNext = nil
		list.length++
	}else {
		bak := list.head
		for bak.pNext != nil {
			bak = bak.pNext  //循环到最后
		}
		bak.pNext = node
		list.length++
	}
}

func (list *SingleLinkList) String() string {
	var listString string
	p := list.head
	for p.pNext != nil {
		listString += fmt.Sprintf("%v-->",p.pNext.value)
		p = p.pNext  //循环
	}
	listString += fmt.Sprintf("nil")
	return listString  //打印链表字符串
}

//中间插入(下一个节点尾部
func (list *SingleLinkList) InsertNodeValueBack(dest interface{},node *SingleLinkNode)  bool{
	phead := list.head
	isfind := false  //是否找到数据
	for phead.pNext !=nil{
		if phead.value == dest { //找到了数据
			isfind = true
			break
		}
		phead = phead.pNext
	}
	if isfind {
		node.pNext = phead.pNext //尾部插入
		phead.pNext = node
		list.length++
		return true
	}else {
		return false
	}
}

//中间插入（下一个节点头部
func (list *SingleLinkList) InsertNodeValueFront(dest interface{},node *SingleLinkNode)  bool{
	phead := list.head
	isfind := false  //是否找到数据
	for phead.pNext !=nil{
		if phead.pNext.value == dest { //找到了数据
			isfind = true
			break
		}
		phead = phead.pNext
	}
	if isfind {
		node.pNext = phead.pNext //
		phead.pNext = node
		list.length++
		return isfind
	}else {
		return isfind
	}
}

//根据index获取节点
func (list *SingleLinkList) GetNodeAtIndex(index int) *SingleLinkNode{
	if index > list.length-1 || index < 0 {
		return nil
	}else {
		phead := list.head
		for index > -1 {
			phead = phead.pNext  //向后循环
			index--  //向后循环过程
		}
		return phead
	}
}

//删除节点
func (list *SingleLinkList) DeleteNode(dest *SingleLinkNode) bool{
	if dest == nil {
		return false
	}
	phead := list.head
	for phead.pNext != nil && phead.pNext != dest {
		phead = phead.pNext
	}
	if phead.pNext == dest {
		phead.pNext = phead.pNext.pNext
		list.length--
		return true
	}else {
		return false
	}
}

//根据index删除节点
func (list *SingleLinkList) Deleteatindex(index int) {
	if index > list.length-1 || index < 0 {
		return
	}else {
		phead := list.head
		for index > -1 {
			phead = phead.pNext  //向后循环
			index--  //向后循环过程
		}
		phead.pNext = phead.pNext.pNext
		list.length--
		return
	}
}

//查找数据
func (list *SingleLinkList) FindString(data string){
	phead := list.head.pNext //指定头部
	for phead.pNext != nil {//循环所有数据
		if strings.Contains(phead.value.(string),data) {//看是否包含
			fmt.Println(phead.value)
		}
		phead = phead.pNext //继续循环
	}
}

//获取链表中间位置（链表的五分之一、三分之一等根据链表长度不一致
func (list *SingleLinkList) GetMid()  *SingleLinkNode {
	if list.head.pNext == nil {
		return nil
	}else {
		phead1 := list.head
		phead2 := list.head
		for phead2 != nil && phead2.pNext != nil{
			phead1 = phead1.pNext
			phead2 = phead2.pNext.pNext  //一次走一步一次走两步
		}
		return phead1  //中间节点
	}
}

//链表反转
func (list *SingleLinkList) ReverseList()  {
	if list.head == nil || list.head.pNext == nil { //链表为空或只有一个节点
		return
	}else {
		//不断循环，原本最后节点的Nil和第一个节点的pnext交换
		var pre *SingleLinkNode  //前一节点
		var cur *SingleLinkNode = list.head.pNext //设置为当前节点
		for cur != nil { //不为nil的时候可以一直循环下去
			curNext := cur.pNext  //后续节点
			cur.pNext = pre   //反转第一步，下一节点地址链接到前一节点

			pre = cur  //持续推进循环
			cur = curNext //持续推进循环
		}
		list.head.pNext.pNext = nil
		list.head.pNext = pre  //头部链接到原来最后一个节点
	}
}
