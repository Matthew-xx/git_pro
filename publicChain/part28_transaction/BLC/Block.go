package BLC

import (
	"bytes"
	"crypto/sha256"
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
	//3. 交易数据(可以打包多笔交易
	Txs []*Transaction
	//4. 时间戳
	Timestamp int64
	//5. Hash
	Hash []byte
	//6. Nonce
	Nonce int64
}

//创建新的区块
func NewBlock(txs []*Transaction,height int64,prevBlockHash []byte) *Block {
	//定义区块
	block := &Block{height,prevBlockHash,txs,time.Now().Unix(),nil,0}

	//调用工作量证明的方法并返回有效的hash和nonce
	pow := NewProofOfWork(block)
	//挖矿验证
	hash,nonce := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	fmt.Println()
	return block
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

//生成创世区块
func CreateGenesisBlock(txs []*Transaction) *Block {
	return NewBlock(txs,1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
}

//将结构体数组(txs)转换成字节数组
func (block *Block) HashTransaction() []byte {
	//将Txs里面的hash全部叠加
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range block.Txs{
		txHashes = append(txHashes,tx.TxHash)
		//bytes.Join([][]byte{tx.TxHash},[]byte{})  //一个key,一个vlue
	}
	txHash = sha256.Sum256(bytes.Join(txHashes,[]byte{}))

	return txHash[:]
}
