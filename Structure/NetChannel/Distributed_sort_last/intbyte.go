package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

//int,string和byte转换

func IntToBytes1(n int) []byte {
	data := int64(n)
	bytebuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytebuffer,binary.BigEndian,data)

	return bytebuffer.Bytes()
}

func BytesToInt1(bts []byte) int {
	bytebuffer := bytes.NewBuffer(bts)
	var data int64
	binary.Read(bytebuffer,binary.BigEndian,&data)

	return int(data)
}

func mainx()  {
	fmt.Println(IntToBytes1(1))
	fmt.Println(BytesToInt1(IntToBytes1(1)))
	fmt.Println(string([]byte("123")))

	//拼接
	myb := IntToBytes1(1)
	myb = append(myb,IntToBytes1(1)...)

	//取其中的一半
	fmt.Println(myb[:len(myb)/2])
}

