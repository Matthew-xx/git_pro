package main

import (
	"fmt"
	"net"
)

func main()  {
	tcpaddr,err := net.ResolveTCPAddr("tcp","127.0.0.1:8848") //注册地址
	if err != nil {
		panic(err)
	}
	conn,err := net.DialTCP("tcp",nil,tcpaddr)  //链接
	if err != nil {
		panic(err)
	}

	conn.Write([]byte("hello mark"))
	buf := make([]byte,1024)
	n,_ := conn.Read(buf)  //读取数据
	fmt.Println(string(buf[:n]))

	for {
		var inputstr string
		fmt.Scanln(&inputstr)
		conn.Write([]byte(inputstr))
		buf := make([]byte,1024)
		n,_ := conn.Read(buf)
		fmt.Println(string(buf[:n]))
	}
}


//0helloword
//1calc  --计算器
