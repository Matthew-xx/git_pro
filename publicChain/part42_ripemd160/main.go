package main

import (
	"fmt"
	"golang.org/x/crypto/ripemd160"
)

func main()  {
	//20个字节，40个数字 ，160位
	//3bae4359ac5fbeaf77091ccf6e1849319efffa2f
	hasher := ripemd160.New()
	hasher.Write([]byte("maxiongxing good"))
	bytes := hasher.Sum(nil)
	fmt.Printf("%x\n",bytes) //16进制
}
