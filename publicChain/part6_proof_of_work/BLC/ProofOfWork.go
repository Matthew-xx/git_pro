package BLC

import (
	"math/big"
)

//256位hash里面前至少有16位为0（设置挖矿难度）
const targetBit = 16

type ProofOfWork struct {
	Block *Block  //要验证的区块
	target *big.Int //大数存储,避免溢出
}

func (proofOfWork *ProofOfWork) Run() ([]byte,int64) {
	return nil,0
}


func NewProofOfWork(block *Block) *ProofOfWork {
	//假设难度为2，一个8位的hash 0000 0001，则需将1左移(8-2)位得到0100 0000 （值为64
	//当证明一个值小于64的hash后也即小于等于( 0010 0000) 便得到了工作量证明


	target := big.NewInt(1)  //创建初始target
	target = target.Lsh(target,256-targetBit) //左移位

	return &ProofOfWork{block,target}
}

