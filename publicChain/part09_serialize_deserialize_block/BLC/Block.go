package BLC

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
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
	//6. Nonce
	Nonce int64
}

//通过block对象调用的方法
/*
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
}*/

//创建新的区块
func NewBlock(data string,height int64,prevBlockHash []byte) *Block {
	//定义区块
	block := &Block{Height:height,
		PrevBlockHash:prevBlockHash,
		Data:[]byte(data),
		Timestamp:time.Now().Unix(),
		Hash:nil,
		Nonce:0,
	}

	//设置hash
	//block.SetHash()

	//调用工作量证明的方法并返回有效的hash和nonce
	pow := NewProofOfWork(block)
	//挖矿验证
	hash,nonce := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	fmt.Println()
	return block
}

//生成创世区块
func CreateGenesisBlock(data string) *Block {
	return NewBlock(data,1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
}

//将区块序列化为字节数组(区块的行为)
func (block *Block) Serialize() []byte {

	var result bytes.Buffer  //定义缓冲
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

//反序列化（谁都可以调用)
func DeserializeBlock(blockBytes []byte) *Block {

	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	err := decoder.Decode(&block)

	if err != nil {
		log.Panic(err)
	}

	return &block
}

