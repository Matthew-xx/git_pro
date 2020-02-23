package main

import (
	"fmt"
	"net"
	"os/exec"
)

//处理器(并发
func MsgHandler(conn net.Conn)  {
	buf := make([]byte,1024)
	defer conn.Close()

	for {
		n,err := conn.Read(buf)
		if err != nil {
			//panic(err)
		}
		//clientip := conn.RemoteAddr()  //远程地址
		if n !=0 {
			//定义协议:当收到的第一个字符为0时，定义为数据。其他为命令
			if string(buf[0:1]) == "0" {
				fmt.Println("client_data:",string(buf[1:n]))

				conn.Write([]byte("收到数据："+string(buf[1:n])+"\n"))
			}else {
				fmt.Println("client_command:",string(buf[1:n]))
				cmd := exec.Command(string(buf[1:n]))  //
				cmd.Run()  //执行命令
				conn.Write([]byte("收到命令："+string(buf[1:n])+"\n"))
			}
		}
	}
}


func main()  {
	server_listener,err := net.Listen("tcp","127.0.0.1:8848")
	if err != nil {
		panic(err)
	}
	defer server_listener.Close()

	for {
		new_conn,err := server_listener.Accept() //接收信息
		if err != nil {
			panic(err)
		}
		go MsgHandler(new_conn)  //处理客户端信息
	}
}



