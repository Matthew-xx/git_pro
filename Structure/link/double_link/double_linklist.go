package double_link

import (
	"fmt"
	"strings"
)

//双链表基本结构
type DoubleLinkList struct {
	head *DoubleLinkNode
	length int
}

//新建一个双链表
func NewDoubleLinkList() *DoubleLinkList {
	head := NewDoubleLinkNode(nil)
	return &DoubleLinkList{head,0}
}
//返回链表长度
func (dlist *DoubleLinkList) GetLength() int {
	return dlist.length
}
//返回第一个节点
func (dlist *DoubleLinkList) GetFirstNode() *DoubleLinkNode {
	return dlist.head.next
}
//头部插入
func (dlist *DoubleLinkList) InsertHead(node *DoubleLinkNode){
	phead := dlist.head
	if phead.next == nil {
		node.next = nil   //phead.next 一样
		phead.next = node //只有一个节点直接连上
		node.prev = phead
		dlist.length++
	}else {
		phead.next.prev = node //标记上一个节点
		node.next = phead.next  //下一个节点

		phead.next = node  //标记头部节点及下一个节点
		node.prev = phead
		dlist.length++
	}
}
//尾部插入
func (dlist *DoubleLinkList) InsertBack(node *DoubleLinkNode){
	phead := dlist.head
	if dlist.head.next == nil {
		node.next = nil   //phead.next 一样
		phead.next = node //只有一个节点直接连上
		node.prev = phead
		dlist.length++
	}else {
		for phead.next != nil {
			phead = phead.next  //循环下去
		}
		phead.next = node  //后缀
		node.prev = phead //前缀
		dlist.length++
	}
}
//展示链表
func (dlist *DoubleLinkList)  String() string {
	var listString1 string
	var listString2 string
	phead := dlist.head
	for phead.next != nil { //正向循环
		listString1 += fmt.Sprintf("%v-->",phead.next.value)
		phead = phead.next
	}
	listString1 += fmt.Sprintf("nil")
	listString1 += "\n"

	for phead != dlist.head { //反向循环
		listString2 += fmt.Sprintf("<--%v",phead.value)
		phead = phead.prev
	}
	listString1 += fmt.Sprintf("nil")

	return listString1+listString2+"\n"  //打印链表字符串
}

//在其他节点前后插入
func (dlist *DoubleLinkList) InsertValueBack(dest *DoubleLinkNode,node *DoubleLinkNode) bool{
	phead := dlist.head
	for phead.next != nil && phead.next != dest {  //循环查找
		phead = phead.next
	}
	if phead.next == dest { //与下一节点相等,则插入
		if phead.next.next != nil {
			phead.next.next.prev = node
		}
		node.next = phead.next.next
		phead.next.next = node
		node.prev = phead.next

		dlist.length++
		return true
	}else {
		return false
	}
}

func (dlist *DoubleLinkList) InsertValueHead(dest *DoubleLinkNode,node *DoubleLinkNode) bool{
	phead := dlist.head
	for phead.next != nil && phead.next != dest {  //循环查找
		phead = phead.next
	}
	if phead.next == dest { //与下一节点相等,则插入
		if phead.next != nil {
			phead.next.prev = node
		}
		node.next = phead.next
		node.prev = phead
		phead.next = node

		dlist.length++
		return true
	}else {
		return false
	}
}

//按数据的位置来实现插入
func (dlist *DoubleLinkList) InsertValueBackByValue(dest interface{},node *DoubleLinkNode) bool{
	phead := dlist.head
	for phead.next != nil && phead.next.value != dest {  //循环查找
		phead = phead.next
	}
	if phead.next.value == dest { //与下一节点相等,则插入
		if phead.next.next != nil {
			phead.next.next.prev = node
		}
		node.next = phead.next.next
		phead.next.next = node
		node.prev = phead.next

		dlist.length++
		return true
	}else {
		return false
	}
}

func (dlist *DoubleLinkList) InsertValueHeadByValue(dest interface{},node *DoubleLinkNode) bool{
	phead := dlist.head
	for phead.next != nil && phead.next.value != dest {  //循环查找
		phead = phead.next
	}
	if phead.next.value == dest { //与下一节点相等,则插入
		if phead.next != nil {
			phead.next.prev = node
		}
		node.next = phead.next
		node.prev = phead
		phead.next = node

		dlist.length++
		return true
	}else {
		return false
	}
}

//根据索引返回节点
func (dlist *DoubleLinkList) GetNodeAtindex(index int) *DoubleLinkNode {
	if index > dlist.length -1 || index<0{
		return nil
	}
	phead := dlist.head
	for index > -1 {
		phead = phead.next
		index--   //计算位置
	}
	return phead
}

//按照节点来删除
func (dlist *DoubleLinkList) DeleteNode(node *DoubleLinkNode) bool{
	if node == nil{
		return false
	}else {
		phead := dlist.head
		for phead.next != nil && phead.next != node {
			phead = phead.next  //循环查找
		}
		if phead.next == node {
			if phead.next.next != nil { //
				phead.next.next.prev = phead  //设置prev
			}
			phead.next = phead.next.next  //设置next ，代表完成了删除
			dlist.length--
			return true
		}else {
			return false
		}

	}
}

//按照索引来删除
func (dlist *DoubleLinkList) DeleteNodeAtindex(index int) bool{
	if index > dlist.length -1 || index<0{
		return false
	}
	phead := dlist.head
	for index > 0 {
		phead = phead.next
		index--   //计算位置
	} //取得索引

	if phead.next.next != nil { //
		phead.next.next.prev = phead  //设置prev
	}
	phead.next = phead.next.next  //设置next ，代表完成了删除
	dlist.length--
	return true
}

func (dlist *DoubleLinkList) FindString(inputstr string)  {
	phead := dlist.head.next
	for phead.next != nil { //正向循环
		if strings.Contains(phead.value.(string),inputstr) { //判断查找
			fmt.Println(phead.value.(string))
		}
		phead = phead.next
	}
}
