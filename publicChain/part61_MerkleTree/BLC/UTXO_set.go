package BLC

import (
	"bytes"
	"encoding/hex"
	"github.com/boltdb/bolt"
	"log"
)

//功能：遍历整个数据库读取所有未花费UTXO，然后将所有的UTXO存储到数据库
//reset重置
//遍历数据库时，返回字典
//存储到数据库
type UTXOSet struct {
	Blockchain *Blockchain
}

const utxoTableName  = "utxotable"

func (utxoSet *UTXOSet) ResetUTXOSet()  {
	err := utxoSet.Blockchain.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(utxoTableName))

		if b!= nil {
			err := tx.DeleteBucket([]byte(utxoTableName))  //删除表
			if err != nil {
				log.Panic(err)
			}
		}

		b,_ = tx.CreateBucket([]byte(utxoTableName)) //不管有没有库都创建
		if b != nil {
			txOutputsMap := utxoSet.Blockchain.FindUTXOMap()  //获取所有未消费输出

			for keyHash,outs := range txOutputsMap{
				txHash,_ := hex.DecodeString(keyHash)
				b.Put(txHash,outs.Serialize())
			}
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

func (utxoSet *UTXOSet) findUTXOForAddress(address string) []*UTXO {
	var utxos []*UTXO

	utxoSet.Blockchain.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(utxoTableName))

		c := b.Cursor()
		//k 是transaction的hash，v 对应outputs对象
		for k,v := c.First(); k!= nil; k,v = c.Next() {
			//fmt.Printf("key=%x , value=%x\n",k,v)
			txOutputs :=DeserializeTXOutputs(v)

			for _,utxo := range txOutputs.UTXOS{
				//判断地址是否匹配（解锁
				if utxo.Output.UnLockScriptPubKeyWithAddress(address) {
					utxos = append(utxos,utxo)
				}
			}
		}
		return nil
	})
	return utxos
}

func (utxoSet *UTXOSet) GetBalance(address string) int64 {
	//找到所有未花费输出
	UTXOS := utxoSet.findUTXOForAddress(address)

	var amount int64
	for _,utxo := range UTXOS{
		amount += utxo.Output.Value
	}

	return amount
}

//返回要凑多少钱，对应的txoutput的tx的hash和index
//某地址查找未打包的交易输出
func (utxoSet *UTXOSet) FindUnPackageSpendableUTXOS(from string,txs []*Transaction) []*UTXO {
	var unUTXOs []*UTXO

	spentTXOutputs := make(map[string][]int)

	//找出已打包的未消费的input对应的数据
	for _,tx := range txs {

		if tx.IsCoinbaseTransaction() == false {
			for _, in := range tx.Vins {
				//是否能够解锁
				publicKeyHash := Base58Decode([]byte(from))

				ripemd160Hash := publicKeyHash[1:len(publicKeyHash) - 4]
				if in.UnLockRipemd160Hash(ripemd160Hash) {

					key := hex.EncodeToString(in.TxHash)

					spentTXOutputs[key] = append(spentTXOutputs[key], in.Vout)
				}

			}
		}
	}

	//找出UTXO
	for _,tx := range txs {

	Work1:
		for index,out := range tx.Vouts {

			if out.UnLockScriptPubKeyWithAddress(from) {
				//fmt.Println("看看是否是俊诚...")
				//fmt.Println(address)

				//fmt.Println(spentTXOutputs)

				if len(spentTXOutputs) == 0 {
					utxo := &UTXO{tx.TxHash, index, out}
					unUTXOs = append(unUTXOs, utxo)
				} else {
					for hash,indexArray := range spentTXOutputs {

						txHashStr := hex.EncodeToString(tx.TxHash)

						if hash == txHashStr {

							var isUnSpentUTXO bool

							for _,outIndex := range indexArray {

								if index == outIndex {
									isUnSpentUTXO = true
									continue Work1
								}

								if isUnSpentUTXO == false {
									utxo := &UTXO{tx.TxHash, index, out}
									unUTXOs = append(unUTXOs, utxo)
								}
							}
						} else {
							utxo := &UTXO{tx.TxHash, index, out}
							unUTXOs = append(unUTXOs, utxo)
						}
					}
				}

			}

		}

	}
	return unUTXOs
}

//
func (utxoSet *UTXOSet) FindSpendableUTXOS(from string,amount int64,txs []*Transaction) (int64,map[string][]int) {
	unPackageUTXOS := utxoSet.FindUnPackageSpendableUTXOS(from,txs)

	//看钱够不够
	//未打包交易中的钱够的话
	spentableUTXO := make(map[string][]int) //存储可用
	var money int64 = 0 //计数器
	for _,UTXO := range unPackageUTXOS{
		money += UTXO.Output.Value
		txHash := hex.EncodeToString(UTXO.TxHash)
		spentableUTXO[txHash] = append(spentableUTXO[txHash],UTXO.Index)

		if money >= amount{
			return money,spentableUTXO
		}
	}

	//如果未打包交易中的钱不够那么需要从表里面取数据
	utxoSet.Blockchain.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(utxoTableName))
		if b != nil {
			c := b.Cursor()
			UTXOBreak:
			//k 是transaction的hash，v 对应outputs对象
			for k,v := c.First(); k!= nil; k,v = c.Next() {
				txOutputs := DeserializeTXOutputs(v)

				for _,utxo := range txOutputs.UTXOS{
					money += utxo.Output.Value
					txHash := hex.EncodeToString(utxo.TxHash)
					spentableUTXO[txHash] = append(spentableUTXO[txHash],utxo.Index)

					if money >= amount {
						break UTXOBreak    //不能直接return  return对应最外面函数
					}
				}
			}
		}
		return nil
	})

	if money < amount{
		log.Panic("余额不足...")
	}

	return money,spentableUTXO
}

//更新（打包完一个区块后就需更新
func (utxoSet *UTXOSet) Update()  {
	//在blocks表中将最新的区块删掉，此时就和utxotable相等

	block := utxoSet.Blockchain.Iterator().Next()  //最新区块

	ins := []*TXInput{}   //暂存需删除的数据
	outsMap := make(map[string]*TXOutputs)  //需保存的数据

	//遍历里面的transaction,找到需删除的数据
	for _,tx := range block.Txs{
		for _,in := range tx.Vins{
			ins = append(ins ,in)
		}  //已花费的输入
	}

	//找出未花费输出
	for _,tx := range block.Txs{
		utxos := []*UTXO{}

		for index,out := range tx.Vouts{

			isSpent := false
			//判断输出有没有被花费
			for _,in := range ins{
				if in.Vout == index && bytes.Compare(tx.TxHash,in.TxHash) == 0 && bytes.Compare(out.Ripemd160Hash,Ripemd160Hash(in.PublicKey)) == 0{
					//输入输出的索引、hash 及公钥吻合,已被花费的out不用理
					isSpent = true
					continue
				}
			}
			if isSpent == false {
				utxo := &UTXO{tx.TxHash,index,out}
				utxos = append(utxos,utxo)
			}
		}

		if len(utxos) > 0 {
			txHash := hex.EncodeToString(tx.TxHash)
			outsMap[txHash] = &TXOutputs{utxos}  //新增到数据库表里面的数据
		}
	}

	err := utxoSet.Blockchain.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(utxoTableName))
		if b != nil {

			//删除
			for _,in := range ins{
				txOutputsBytes := b.Get(in.TxHash) //通过hash匹配拿到表里面的数据

				if len(txOutputsBytes) == 0 {
					continue
				}
				txOutputs := DeserializeTXOutputs(txOutputsBytes)

				UTXOS := []*UTXO{}
				isNeedDelete := false  //判断是否需要删除

				for _,utxo := range txOutputs.UTXOS{
					if in.Vout == utxo.Index {
						if bytes.Compare(utxo.Output.Ripemd160Hash,Ripemd160Hash(in.PublicKey)) == 0{
							//需要删除的交易的公钥与未花费输出的数据比对是否相等
							isNeedDelete = true
						}else {
							UTXOS = append(UTXOS,utxo)
						}
					}
				}

				if isNeedDelete {
					b.Delete(in.TxHash)  //删除hash所对应的键值对

					if len(UTXOS) >0 {

						preTXOutputs := outsMap[hex.EncodeToString(in.TxHash)]
						preTXOutputs.UTXOS = append(preTXOutputs.UTXOS,UTXOS...)
						outsMap[hex.EncodeToString(in.TxHash)] = preTXOutputs
					}
				}
			}

			//新增
			for keyHash,outPuts := range outsMap {
				keyHashBytes,_ := hex.DecodeString(keyHash)  //将hash转换成字节数组
				b.Put(keyHashBytes,outPuts.Serialize())
			}
		}


		return nil
	})
	if err != nil {
		log.Panic(err)
	}

}
