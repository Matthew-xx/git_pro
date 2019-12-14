package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//256位hash里面前至少有16位(二进制为4位）为0（设置挖矿难度）
const targetBit = 16

//数据拼接，返回字节数组
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.HashTransaction(),
			IntToHex(pow.Block.Timestamp),
			IntToHex(int64(targetBit)),
			IntToHex(int64(nonce)),
			IntToHex(int64(pow.Block.Height)),
		}, []byte{},
		)
	return data
}

type ProofOfWork struct {
	Block *Block  //要验证的区块
	target *big.Int //大数存储,避免溢出
}

func (proofOfWork *ProofOfWork) Run() ([]byte,int64) {
	//1. 将block的属性拼接成字节数组
	//2. 生成hash
	//3. 判断hash有效性，若满足条件，跳出循环
	nonce := 0
	var hash [32]byte
	var hashInt big.Int  //存储新生成的hash

	for {
		dataBytes := proofOfWork.prepareData(nonce) //准备数据
		hash = sha256.Sum256(dataBytes) //生成hash
		fmt.Printf("\r%x",hash)
		hashInt.SetBytes(hash[:]) //存储hash

		//判断hashInt是否小于block里面的target（等于1代表target大于后面的hashInt
		//也即挖矿难度得到了证明
		if proofOfWork.target.Cmp(&hashInt) == 1{
			break;
		}
		nonce = nonce +1
	}
	
	return hash[:],int64(nonce)
}


func NewProofOfWork(block *Block) *ProofOfWork {
	//假设难度为2，一个8位的hash 0000 0001，则需将1左移(8-2)位得到0100 0000 （值为64
	//当证明一个值小于64的hash后也即小于等于( 0010 0000) 便得到了工作量证明


	target := big.NewInt(1)  //创建初始target
	target = target.Lsh(target,256-targetBit) //左移位

	return &ProofOfWork{block,target}
}

