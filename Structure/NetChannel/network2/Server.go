package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"time"
)

//处理器(并发
func MsgHandler(conn net.Conn)  {
	buf := make([]byte,1024)
	defer conn.Close()

	for {
		n,err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn close")
			return
			//panic(err)
		}
		//clientip := conn.RemoteAddr()  //远程地址
		msg := buf[1:n]
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

		beatch := make(chan byte)
		go HeartBeat(conn,beatch,30)
		go HeartChanHandler(msg,beatch)
	}
}
//判断一定时间内有没有产生通信
func HeartBeat(conn net.Conn,heartchan chan byte,timeout int)  {
	select {
	case hc:= <- heartchan:
		fmt.Println(string(hc))
		log.Println("heartchan",string(hc))
		//产生通信，计时器归零
		conn.SetDeadline(time.Now().Add(time.Duration(timeout)*time.Second))
	case <- time.After(time.Second*30):
		fmt.Println("time out",conn.RemoteAddr())  //客户端访问超时
		log.Println("time out",conn.RemoteAddr())
		conn.Close()   //30秒内无访问，退出
	}
}

//处理心跳的channel
func HeartChanHandler(n[]byte,beatch chan byte)  {
	for _,v := range n{
		beatch <- v   //压入数据
	}
	close(beatch)  //关闭管道
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



