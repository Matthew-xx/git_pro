package main

import (
	"fmt"
	"math"
)

//锦标赛排序
//堆排序是将所有数放树上，锦标赛只把部分数据放树上
//二叉树上两两比较选出最小值，若有某个子节点为空则用无穷大替代。
//当某个子节点的数被选出为最小值后，用无穷大替代该子节点，再进行下一轮选最小值。
//将每轮选出的最小值依次放进存储容器中

type node struct {
	value int   //叶子的数据
	isok  bool  //叶子状态是否无穷大
	rank  int   //叶子的排序
}

func pow(x,y int) int {  //求x的y次方
	return int(math.Pow(float64(x),float64(y)))
}

func compileAndUp(tree *[]node,leftnode int)  {
	rightnode := leftnode +1

	if !(*tree)[leftnode].isok || ((*tree)[rightnode].isok && (*tree)[leftnode].value > (*tree)[rightnode].value) {
		mid := (leftnode-1)/2  //中间节点（父节点）
		(*tree)[mid] = (*tree)[rightnode]  //将较小的值复制的父子节点
	}else {
		mid := (leftnode-1)/2
		(*tree)[mid] = (*tree)[leftnode]
	}
}

func TreeSelectSort(arr []int) []int {
	var level int//树的层数
	var result = make([]int,0,len(arr)) //保存最终结果
	for pow(2,level) < len(arr) { //求出可以覆盖所有元素的层数
		level++
	}
	var leaf = pow(2,level) //叶子的数量
	var tree = make([]node,leaf*2-1) //构造树节点数量
	//填充叶子节点
	for i:=0;i<len(arr);i++ {
		tree[leaf+i-1] = node{arr[i],true,i}
	}

	//进行对比，选出两者中最小值
	for i:=0;i<level;i++ {
		nodecount := pow(2,level-i)  //每次处理则降低一个层级
		for j:=0;j<nodecount/2;j++ {
			leftnode := nodecount -1+j*2
			compileAndUp(&tree,leftnode)
		}

	}
	result = append(result,tree[0].value)  //保存最顶端的最小值

	//选出第一个后，还有n-1个循环
	for t:=0;t<len(arr)-1;t++ {
		winnode := tree[0].rank+leaf-1  //记录赢得的节点（上次选出最小值的点
		tree[winnode].isok = false  //理解为修改为无穷大
		for i:=0;i<level;i++ {
			leftNode := winnode
			if winnode %2 == 0{  //处理奇偶节点数
				leftNode = winnode -1
			}
			compileAndUp(&tree,leftNode)

			winnode = (leftNode-1)/2  //保存中间节点（父节点
		}
		result = append(result,tree[0].value)
		fmt.Println(result)
	}
	return result
}

func main()  {
	var length = 10
	var mymap = make(map[int]int,length)
	var obj []int
	//构造map，随机数据
	for i:=0;i<length;i++ {
		mymap[i] = i
	}
	for k,_ := range mymap{
		obj = append(obj,k) //叠加
	}
	fmt.Println(obj)
	fmt.Println(TreeSelectSort(obj))
}
