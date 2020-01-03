package BLC

import "crypto/sha256"

type MerkleTree struct {
	RootNode *MerkleNode
}

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

//merklenode{nil,nil,tx1bytes} , merklenode{nil,nil,tx2bytes}

//merklenode{merklenode{nil,nil,tx1bytes},merklenode{nil,nil,tx2bytes},sha256(tx1bytes,tx2bytes)}

//merklenode{
//   left:merklenode{merklenode{nil,nil,tx1bytes},merklenode{nil,nil,tx2bytes},sha256(tx1bytes,tx2bytes)}
//   right: ...
//   sha256(sha256(tx1bytes,tx2bytes) + ...)
//   }


func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode

	if len(data)%2 != 0 {
		data = append(data,data[len(data)-1])   //奇数复制一份
	}

	//创建叶子节点
	for _,dataum := range data{
		node := NewMerkleNode(nil,nil,dataum)  //dataum对应transaction序列化后的字节数组
		nodes = append(nodes,*node)
	}

	for i:= 0;i<len(data)/2;i++ {
		var newLevel []MerkleNode
		for j :=0;j<len(nodes);j+=2 {
			node := NewMerkleNode(&nodes[j],&nodes[j+1],nil)
			newLevel = append(newLevel,*node)
		}

		if len(newLevel) %2 !=0 {
			newLevel = append(newLevel,newLevel[len(newLevel)-1])
		}
		nodes = newLevel
	}

	mTree := MerkleTree{&nodes[0]}
	return &mTree
}

func NewMerkleNode(left,right *MerkleNode,data []byte) *MerkleNode {
	mNode := MerkleNode{}

	//创建叶子节点
	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		mNode.Data = hash[:]
	}else { //非叶子节点
		prevHashes := append(left.Data,right.Data...)
		hash := sha256.Sum256(prevHashes)
		mNode.Data = hash[:]
	}

	mNode.Left = left
	mNode.Right = right

	return &mNode
}

