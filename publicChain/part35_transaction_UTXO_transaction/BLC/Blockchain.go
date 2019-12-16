package BLC

import (
	"encoding/hex"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"
)

const dbName  = "blockchain.db"
const blockTableName  = "blocks"

//创建区块链结构体
type Blockchain struct {
	Tip []byte    //存储最新区块的hash值
	DB  *bolt.DB    //数据库
	//Blocks []*Block    //存储有序的区块
}

//迭代器
func (blockchain *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{blockchain.Tip,blockchain.DB}
}
//判断是否存在
func DbExists() bool {
	if _,err := os.Stat(dbName); os.IsNotExist(err) {
		return false
	}

	return true
}

//遍历输出所有区块的信息

func (blc *Blockchain) PrintChain() {
	blockchainIterator := blc.Iterator()

	for  {
		block := blockchainIterator.Next()
		fmt.Println("---------------------------------------")
		fmt.Printf("Height: %d\n",block.Height)
		fmt.Printf("prevHash: %x\n",block.PrevBlockHash)
		//fmt.Printf("Txs: %v\n",block.Txs)
		fmt.Printf("Timestamp: %s\n",time.Unix(block.Timestamp,0).Format("2006-01-02 03:04:25 PM"))
		fmt.Printf("Hash: %x\n",block.Hash)
		fmt.Printf("Nonce: %d\n",block.Nonce)
		fmt.Println("Txs:")
		for _,tx := range block.Txs{
			//遍历交易数组
			fmt.Printf("%x\n",tx.TxHash)
			fmt.Println("Vins:")
			for _,in := range tx.Vins{
				//要消费“人”的具体信息
				fmt.Printf("%x\n",in.Txid)
				fmt.Printf("%d\n",in.Vout)
				fmt.Printf("%s\n",in.ScriptSig)
			}
			fmt.Println("Vout:")
			for _,out := range tx.Vouts{
				//收款人的具体信息
				fmt.Printf("%d\n",out.Value)
				fmt.Printf("%s\n",out.ScriptPubKey)
			}
		}

		//判断结束
		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)

		if big.NewInt(0).Cmp(&hashInt) == 0{
			break;
		}
	}
}

//增加区块到区块链里面
func (blc *Blockchain) AddBlockToBlockchain(txs []*Transaction)  {

	err := blc.DB.Update(func(tx *bolt.Tx) error {
		//获取表
		b := tx.Bucket([]byte(blockTableName))

		if b != nil {
			//获取最新区块并反序列化
			blockBytes := b.Get(blc.Tip)
			block := DeserializeBlock(blockBytes)

			//创建新区块(高度加1
			newBlock := NewBlock(txs, block.Height+1, block.Hash)

			//将区块序列化并存储到数据库
			err := b.Put(newBlock.Hash,newBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			//更新数据库里面的“l“对应的hash
			err = b.Put([]byte("l"),newBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			//更新blockchain中的tip
			blc.Tip = newBlock.Hash
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

//创建带有创世区块的区块链
//改成只需要创建创世区块并存储到数据库，而不必要返回对象
func CreateBlockchainWithGenesisBlock(address string) *Blockchain{
	//判断数据库是否存在
	if DbExists() {
		fmt.Println("创世区块已经存在")
		os.Exit(1)
	}

	fmt.Println("正在创建创世区块...")
	//创建或打开数据库
	db,err := bolt.Open(dbName,0600,nil)
	if err != nil {
		log.Fatal(err)
	}

	var genesisHash []byte

	//更新
	err = db.Update(func(tx *bolt.Tx) error {
		//获取表（创建

		b ,err := tx.CreateBucket([]byte(blockTableName))
		if err != nil {
			log.Panic(err)
		}

		if b != nil {
			//创建了一个coinbase交易
			txCoinbase := NewCoinbaseTransaction(address)

			//传入transaction数组（之前是字节数组
			genesisBlock := CreateGenesisBlock([]*Transaction{txCoinbase})
			//将创世区块(序列化后）存储到库里
			err := b.Put(genesisBlock.Hash,genesisBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			// 存储最新区块的hash
			err = b.Put([]byte("l"),genesisBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			genesisHash = genesisBlock.Hash
		}
		return nil
	})
	return &Blockchain{genesisHash,db}
}

//返回blockchain
func GetBlockChainObject() *Blockchain {

	db,err := bolt.Open(dbName,0600,nil)
	if err != nil {
		log.Fatal(err)
	}

	var tip []byte
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			tip = b.Get([]byte("l"))  //读取最新区块hash

		}
		return nil
	})

	return &Blockchain{tip,db}
}

//挖掘新区块
func (blockchain *Blockchain)MineNewBlock(from []string,to []string,amount []string) {
	//fmt.Println(from)
	//fmt.Println(to)
	//fmt.Println(amount)

	amount_value , _ := strconv.Atoi(amount[0])
	tx := NewSimpleTransaction(from[0],to[0],amount_value)


	//通过相关算法建立transaction数组
	var txs []*Transaction
	txs = append(txs,tx)

	//获取最新区块的相关信息
	var block *Block
	blockchain.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			hash := b.Get([]byte("l")) //通过l 拿到最新的hash
			blockBytes := b.Get(hash)  //通过hash拿到区块（字节数组
			block = DeserializeBlock(blockBytes) //反序列化拿到最新的区块
		}
		return nil
	})

	//建立新的区块
	//并不是调用(新建)函数创建新区块，而是直接从数据库中读取上一区块的信息
	block = NewBlock(txs,block.Height+1,block.Hash) //交易数组，上一区块高度+1，上一区块hash

	//将新区块存储到数据库
	blockchain.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))

		if b != nil {
			b.Put(block.Hash,block.Serialize())
			b.Put([]byte("l"),block.Hash)

			blockchain.Tip = block.Hash
		}
		return nil
	})
}

//如果一个地址对应的TxOutput未花费，那么将这个transaction添加到数组中返回未消费输出
func (blockchain *Blockchain) UnUTXOs(address string) []*TxOutput {
	//遍历数据库查找

	var unUTXOs []*TxOutput //用数组存储未消费输出
	spentTxOutputs := make(map[string][]int)   //存储某用户的所有输入（已消费的

	blockIterator := blockchain.Iterator()

	for  {
		block := blockIterator.Next()  //从迭代器中拿到区块

		//fmt.Println(block)
		//fmt.Println()  //空行

		for _,tx := range block.Txs{
			//遍历每个block里面的transaction(txHash,Vins,Vouts)
			//判断输入Vins (创世区块除外

			if tx.IsCoinbaseTransaction() == false { //除创世区块外的其他交易
				for _,in := range tx.Vins{
					//是否能够解锁(用户姓名是否相等)
					if in.UnlockScriptSigWithAddress(address) {
						//如果为真，则说明是所取的用户
						//存储，id做key,
						key := hex.EncodeToString(in.Txid)
						spentTxOutputs[key] = append(spentTxOutputs[key],in.Vout) //
					}
				}
			}
			
			//判断输出Vouts
			for index,out := range tx.Vouts {
				//解锁
				if out.UnlockScriptPubKeyWithAddress(address) {
					fmt.Println(out)
					//用户对应一致后要判断输出是否已被消费
					//spentTxOutputs已经存储了所有已消费的,
					if spentTxOutputs != nil{

						if len(spentTxOutputs) != 0 { //某个用户只有收款，在上面的付款数组中没记录到
							for txHash,indexArray := range spentTxOutputs{

								for _,i := range indexArray{
									if index == i && txHash == hex.EncodeToString(tx.TxHash) {
										//说明这笔钱已被花费
										continue
									} else {
										unUTXOs = append(unUTXOs,out)
									}
								}
							}
						}else {
							unUTXOs = append(unUTXOs,out)
						}
					}
				}
			}
		}

		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)

		if hashInt.Cmp(big.NewInt(0)) == 0{
			break;  //上面两个相等，说明已经遍历到创世区块了
		}
	}
	return unUTXOs
}

