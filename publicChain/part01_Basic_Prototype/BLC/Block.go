package BLC

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	//1. 区块高度
	Height int64
	//2. 上一区块hash
	PrevBlockHash []byte
	//3. 交易数据
	Data []byte
	//4. 时间戳
	Timestamp int64
	//5. Hash
	Hash []byte
}

//通过block对象调用的方法
func (block *Block) SetHash() {
	//1. height ,时间戳转成字节数组 []byte
	heightBytes := IntToHex(block.Height)

	timeString := strconv.FormatInt(block.Timestamp,2) //2进制
	timeBytes := []byte(timeString)
	//2. 拼接所有属性
	blockBytes := bytes.Join([][]byte{heightBytes,block.PrevBlockHash,block.Data,timeBytes,block.Hash},[]byte{})
	//3.生成hash
	hash := sha256.Sum256(blockBytes)

	block.Hash = hash[:]  //因为hash是32字节，但block.Hash是切片数组
}

//创建新的区块
func NewBlock(data string,height int64,prevBlockHash []byte) *Block {
	//定义区块
	block := &Block{Height:height,
		PrevBlockHash:prevBlockHash,
		Data:[]byte(data),
		Timestamp:time.Now().Unix(),
		Hash:nil,
	}

	//设置hash
	block.SetHash()

	return block
}