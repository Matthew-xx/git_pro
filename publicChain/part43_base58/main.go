package main

import (
	"../part43_base58/BLC"
	"fmt"
)


func main()  {
	//20个字节，40个数字 ，160位
	//3bae4359ac5fbeaf77091ccf6e1849319efffa2f

	bytes := []byte("hello my love")
	//bytes的长度不一样产生的bytes58长度也不一样
	bytes58 := BLC.Base58Encode(bytes)  //加密
	fmt.Printf("%x\n",bytes58)
	fmt.Printf("%s\n",bytes58)  //地址

	bytesStr := BLC.Base58Decode(bytes58) // 解密
	fmt.Printf("%s\n",bytesStr[1:]) //因为加密的时候在字符串前面加了东西
}
