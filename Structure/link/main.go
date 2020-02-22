package main

import (
	"./double_link"
	"./linking"
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main1()  {
	list := linking.NewSingleLinkList()
	node1 := linking.NewSingleLinkNode(1)
	node2 := linking.NewSingleLinkNode(2)
	node3 := linking.NewSingleLinkNode(3)
	list.InsertNodeBack(node1)
	fmt.Println(list)
	list.InsertNodeBack(node2)
	fmt.Println(list.GetFirstNode())
	list.InsertNodeBack(node3)
	fmt.Println(list)
	node4 := linking.NewSingleLinkNode(4)
	list.InsertNodeValueFront(2,node4)
	fmt.Println(list)
	fmt.Println(list.GetMid().Value())
	list.ReverseList()
	fmt.Println(list)
}

//单链表读取数据
func main2()  {
	list := linking.NewSingleLinkList()
	path := "F:\\Software\\go_path\\my_pro\\Structure\\data\\datax.txt"
	file, _ := os.Open(path)
	defer file.Close()

	br := bufio.NewReader(file)
	i := 0  //统计共多少行
	for  {
		line,_,end := br.ReadLine()  //逐行读取
		if end == io.EOF { //文件关闭
			break
		}

		linestr := string(line)
		nodestr := linking.NewSingleLinkNode(linestr)
		list.InsertNodeFront(nodestr)  //将数据插入链表
		i++
	}
	//fmt.Println(list)
	fmt.Println(list.GetNodeAtIndex(6))
	for ;; {
		fmt.Println("要搜索的数据：")
		var QQ string
		fmt.Scanln(&QQ)  //查询QQ
		startTime := time.Now()

		list.FindString(QQ)
		fmt.Println("搜索耗时：",time.Since(startTime))
	}
}

func main3()  {
	dlist := double_link.NewDoubleLinkList()
	node1 := double_link.NewDoubleLinkNode(1)
	node2 := double_link.NewDoubleLinkNode(2)
	node3 := double_link.NewDoubleLinkNode(3)
	node4 := double_link.NewDoubleLinkNode(4)
	node5 := double_link.NewDoubleLinkNode(5)
	node6 := double_link.NewDoubleLinkNode(6)
	//node7 := double_link.NewDoubleLinkNode(7)
	node8 := double_link.NewDoubleLinkNode(7)
	dlist.InsertHead(node1)
	dlist.InsertHead(node2)
	dlist.InsertHead(node3)
	dlist.InsertHead(node4)
	dlist.InsertHead(node5)
	dlist.InsertHead(node6)
	fmt.Println(dlist.String())
	//dlist.InsertValueBack(node3,node7)
	dlist.InsertValueBackByValue(2,node8)
	fmt.Println(dlist.String())
	//dlist.DeleteNode(node3)
	dlist.DeleteNodeAtindex(2)
	fmt.Println(dlist.String())
}

//用双链表读取数据
func main()  {
	pathlist := []string{"F:\\Software\\go_path\\my_pro\\Structure\\data\\datax.txt",
		"F:\\Software\\go_path\\my_pro\\Structure\\data\\datay.txt",}
	dlist := double_link.NewDoubleLinkList()
	for i:=0;i<len(pathlist);i++ {
		path := pathlist[i]
		file, _ := os.Open(path)
		//defer file.Close()

		br := bufio.NewReader(file)
		for  {
			line,_,end := br.ReadLine()  //逐行读取
			if end == io.EOF { //文件关闭
				break
			}
			linestr := string(line)
			node := double_link.NewDoubleLinkNode(linestr) //新建节点
			dlist.InsertHead(node)  //插入节点
		}
	}
	fmt.Println("载入完成:",dlist.GetLength())
	for ;; {
		fmt.Println("要搜索的数据：")
		var QQ string
		fmt.Scanln(&QQ)  //查询QQ
		startTime := time.Now()

		dlist.FindString(QQ)
		fmt.Println("搜索耗时：",time.Since(startTime))
	}
}



