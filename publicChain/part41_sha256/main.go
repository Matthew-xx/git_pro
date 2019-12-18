package main

import (
	"crypto/sha256"
	"fmt"
)

func main()  {
	//生成32个字节，64个数字，256位的hash
	//ea b95f96f5492e7a56c8cea5cd4a0df39d1e8872c875d53811568ad91de83ff5
	//两个数字为8位字节
	hasher := sha256.New()
	hasher.Write([]byte("maxiongxing good"))
	bytes := hasher.Sum(nil)
	fmt.Printf("%x\n",bytes) //16进制
}
