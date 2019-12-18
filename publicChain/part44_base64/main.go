package main

import (
	"encoding/base64"
	"fmt"
)

func main()  {
	msg := "hello my love"
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println(encoded)

	decoded,err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("decode error:",err)
		return
	}
	fmt.Println(string(decoded))
}