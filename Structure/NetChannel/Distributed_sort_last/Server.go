package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)
func IntToBytesx(n int) []byte {
	data := int64(n)
	bytebuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytebuffer,binary.BigEndian,data)

	return bytebuffer.Bytes()
}
func BytesToIntx(bts []byte) int {
	bytebuffer := bytes.NewBuffer(bts)
	var data int64
	binary.Read(bytebuffer,binary.BigEndian,&data)

	return int(data)
}

func ServerMsgHandler(conn net.Conn){
	buf := make([]byte,16)
	defer conn.Close()

	arr := []int{}  //数组保存数据
	for {
		n,err := conn.Read(buf)
		if err != nil {
			fmt.Println("server close")
			return
		}
		if n == 16{
			data1 := BytesToIntx(buf[:len(buf)/2])
			data2 := BytesToIntx(buf[len(buf)/2:])
			if data1 == 0 && data2 == 0{  //0 0 开始接收数据
				arr = make([]int,0,0)
			}
			if data1 == 1{
				arr = append(arr,data2)
			}
			if data1 == 0 && data2 == 1{  //0 1 结束接收数据
				fmt.Println("数组接收完成:",arr)

				arr = make([]int,0,0)  //开辟数据准备下次接收
			}
		}

	}

}
func main()  {
	tcpaddr,err := net.ResolveTCPAddr("tcp","127.0.0.1:8848") //注册地址
	if err != nil {
		panic(err)
	}
	conn,err := net.DialTCP("tcp",nil,tcpaddr)  //链接
	if err != nil {
		panic(err)
	}
//0 0 代表开始传输
//1 1
//1 2 代表传输的数据
//0 1代表结束传输
	go ServerMsgHandler(conn)  //双工
	arr := []int{1,4,9,2,6,7,3,10,8,5,18,24}
	legnth := len(arr)

	mybstart := IntToBytesx(0)
	mybstart = append(mybstart,IntToBytesx(0)...)
	conn.Write(mybstart)

	for i:=0;i<legnth;i++ {
		mybdata := IntToBytesx(1)
		mybdata = append(mybdata,IntToBytesx(arr[i])...)
		conn.Write(mybdata)
	}

	mybend := IntToBytesx(0)
	mybend = append(mybend,IntToBytesx(1)...)
	conn.Write(mybend)

	time.Sleep(time.Second*30)
}


//0helloword
//1calc  --计算器
