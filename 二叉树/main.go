// 二叉树 project main.go
package main

import (
	"btree"
	"fmt"
)

func main() {
	//创建二叉树
	root := btree.NewNode(nil, nil)
	root.SetData("root node")
	a := btree.NewNode(nil, nil)
	a.SetData("left node")
	al := btree.NewNode(nil, nil)
	al.SetData(100)
	ar := btree.NewNode(nil, nil)
	ar.SetData(3.14)
	a.Left = al
	a.Right = ar
	b := btree.NewNode(nil, nil)
	b.SetData("right node")
	root.Left = a
	root.Right = b
	// 使用 Operater 接口实现对二叉树的基本操作
	var it btree.Operater
	it = root
	it.PrintBT()
	fmt.Println()
	fmt.Println("The depths of the Btree is:", it.Depth())
	fmt.Println("The leaf counts of the Btree is:", it.LeafCount())
	// 使用 Order 接口实现对二叉树的基本操作
	var it2 btree.Order
	it2 = root
	it2.PreOrder() //先序遍历
	fmt.Println()
	it2.InOrder() //中序遍历
	fmt.Println()
	it2.PostOrder() //后序遍历

}
