package main

import (
	"../part43_base58/BLC"
	"crypto/sha256"
	"fmt"
)


func main()  {
	//20个字节，40个数字 ，160位
	//3bae4359ac5fbeaf77091ccf6e1849319efffa2f

	bytes := []byte("hello my love")

	hasher := sha256.New()
	hasher.Write(bytes)
	hash := hasher.Sum(nil)
	//此时长度不变
	bytes58 := BLC.Base58Encode(hash)  //加密
	fmt.Printf("%x\n",bytes58)
	fmt.Printf("%s\n",bytes58)  //地址

	bytesStr := BLC.Base58Decode(bytes58) // 解密
	fmt.Printf("%x\n",bytesStr[1:]) //因为加密的时候在字符串前面加了东西
}

/*
公钥 ——》pubkeyhash: ripemd160(sha256(公钥)) :20个字节 ——》
address = base58（version:0x00 + pubkeyhash + sha256(sha256(pubkeyhash)) :截取最末尾4个字节 ）
共25字节 1KoKTUHRrGn3ZChS1zPyzn721oiZzJc84V
 */
/*
1、创建钱包{私钥、公钥}
2、先将公钥进行一次sha256,再进行一次160hash（RIPEMD160），返回20字节数组
3、base58编码(25字节)
 */
