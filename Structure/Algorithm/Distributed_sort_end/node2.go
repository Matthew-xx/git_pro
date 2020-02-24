package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"sort"
	"time"
)

func BytesToInt1(bts []byte) int {
	bytebuffer := bytes.NewBuffer(bts)
	var data int64
	binary.Read(bytebuffer,binary.BigEndian,&data)

	return int(data)
}
func IntToBytes1(n int) []byte {
	data := int64(n)
	bytebuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytebuffer,binary.BigEndian,data)

	return bytebuffer.Bytes()
}

//处理器(并发
func MsgHandler1(conn net.Conn)  {
	buf := make([]byte,16)
	defer conn.Close()

	arr := []int{}  //数组保存数据
	for {
		n,err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn close")
			return
		}
		if n == 16{
			data1 := BytesToInt1(buf[:len(buf)/2])
			data2 := BytesToInt1(buf[len(buf)/2:])
			if data1 == 0 && data2 == 0{  //0 0 开始接收数据
				arr = make([]int,0,0)
			}
			if data1 == 1{
				arr = append(arr,data2)
			}
			if data1 == 0 && data2 == 1{  //0 1 结束接收数据
				fmt.Println("数组接收完成:",arr)
				sort.Ints(arr)
				fmt.Println("数组排序完成:",arr)

				//写入
				mybstart := IntToBytes1(0)
				mybstart = append(mybstart,IntToBytes1(0)...)
				conn.Write(mybstart)

				for i:=0;i<len(arr);i++ {
					mybdata := IntToBytes1(1)
					mybdata = append(mybdata,IntToBytes1(arr[i])...)
					conn.Write(mybdata)
				}

				mybend := IntToBytes1(0)
				mybend = append(mybend,IntToBytes1(1)...)
				conn.Write(mybend)

				arr = make([]int,0,0)
			}
		}

	}
}
//判断一定时间内有没有产生通信
func HeartBeat1(conn net.Conn,heartchan chan byte,timeout int)  {
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
func HeartChanHandler1(n[]byte,beatch chan byte)  {
	for _,v := range n{
		beatch <- v   //压入数据
	}
	close(beatch)  //关闭管道
}

func main()  {
	server_listener,err := net.Listen("tcp","127.0.0.1:7001")
	if err != nil {
		panic(err)
	}
	defer server_listener.Close()

	for {
		new_conn,err := server_listener.Accept() //接收信息
		if err != nil {
			panic(err)
		}
		go MsgHandler1(new_conn)  //处理客户端信息
	}
}



