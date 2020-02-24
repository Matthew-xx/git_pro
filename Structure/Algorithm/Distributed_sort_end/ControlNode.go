package main

import (
	"./pipelineMiddleware"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
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

func ServerMsgHandler(conn net.Conn) <- chan int{
	out := make(chan int,1024)
	buf := make([]byte,16)
	//defer conn.Close()

	arr := []int{}  //数组保存数据
	for {
		n,err := conn.Read(buf)
		if err != nil {
			fmt.Println("server close")
			return nil
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
				for i:=0; i<len(arr);i++ {
					out <- arr[i]   //数组压入管道
				}
				close(out)
				return out
				arr = make([]int,0,0)  //开辟数据准备下次接收
			}
		}

	}
	return nil
}

func SendArray(arr []int,conn net.Conn)  {
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
}


func main()  {
	arrlist := [][]int{{1,4,9,2,6,7,3,10,8,5,18,24},{11,14,19,12,16,17,13,20,28,15,34,55}}
	sortresults := [] <- chan int {}

	for i:=0; i<2; i++ {
		tcpaddr,err := net.ResolveTCPAddr("tcp","127.0.0.1:700"+strconv.Itoa(1+i)) //注册地址
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
		SendArray(arrlist[i],conn)
		sortresults = append(sortresults,ServerMsgHandler(conn)) //双工

	}
	lastout := pipelineMiddleware.Merge(sortresults[0],sortresults[1])  //两两归并
	for v:= range lastout{
		fmt.Printf("%d ",v)
	}
	time.Sleep(time.Second*30)
}
