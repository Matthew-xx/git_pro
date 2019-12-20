package BLC

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//存储多个wallet

const walletfile  = "Wallet.dat"

type Wallets struct {
	WalletsMap map[string]*Wallet
}

//map[key:obj] :存储着key（地址）和object（对象）
//创建钱包集合
func NewWallets() (*Wallets,error) {
	//先判断文件是否存在，不存在创建新钱包
	if _,err := os.Stat(walletfile); os.IsNotExist(err) {
		wallets := &Wallets{}
		wallets.WalletsMap = make(map[string]*Wallet)
		return wallets,err
	}

	fileContent,err := ioutil.ReadFile(walletfile)  //文件存在，读取文件
	if err != nil {
		log.Panic(err)
	}

	var wallets Wallets  //读取文件返回wallets对象
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	//反序列化（文件存储的并不是字节数组，先转成字节数组再反序列化后才是原来存储的字节数组
	err = decoder.Decode(&wallets)
	if err != nil {
		log.Panic(err)
	}

	return &wallets,nil
}

//创建新钱包
func (w *Wallets) CreateNewWallet()  {
	wallet := NewWallet()
	fmt.Printf("Address is: %s\n",wallet.GetAddress())
	w.WalletsMap[string(wallet.GetAddress())] = wallet

	w.SaveWallets()  //创建钱包后存储
}

/*  //根据地址获取钱包对象
func (ws Wallets) GetWallet(address string) Wallet {
	return *ws.WalletsMap[address]
}

//加载钱包文件
func (ws *Wallets) LoadFromFile() error {
	//判断文件是否存在
	if _,err := os.Stat(walletfile); os.IsNotExist(err) {
		return err
	}

	fileContent,err := ioutil.ReadFile(walletfile)
	if err != nil {
		log.Panic(err)
	}

	var wallets Wallets
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	//反序列化（文件存储的并不是字节数组，先转成字节数组再反序列化后才是原来存储的字节数组
	err = decoder.Decode(&wallets)
	if err != nil {
		log.Panic(err)
	}

	ws.WalletsMap = wallets.WalletsMap  //序列化后的数组覆盖新wallets（即多条覆盖0条）

	return nil
}*/

//把wallets信息保存到一个文件中
func (w *Wallets) SaveWallets()  {
	var content bytes.Buffer

	gob.Register(elliptic.P256()) //Register注册,是为了可以序列化任何类型（包括接口

	encoder := gob.NewEncoder(&content)  //序列化
	err := encoder.Encode(&w)
	if err != nil {
		log.Panic(err)
	}

	//将序列化后的数据写入文件，原来文件的数据会被覆盖
	err = ioutil.WriteFile(walletfile,content.Bytes(),0644)  //0644表读写权限
	if err != nil {
		log.Panic(err)
	}
}

