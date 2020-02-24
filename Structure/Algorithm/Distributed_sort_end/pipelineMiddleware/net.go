package pipelineMiddleware

import (
	"bufio"
	"net"
)

//通过IP地址的网络写入数据
func NetWorkWrite(addr string,in <- chan int)  {
	listen,err := net.Listen("tcp",addr)  //监听
	if err != nil {
		panic(err)
	}
	go func() {
		defer listen.Close() //关闭网络
		conn,err := listen.Accept() //接收信息
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		writer := bufio.NewWriter(conn)  //写入数据
		defer writer.Flush()
		WriteSlink(writer,in)
	}()
}

//在ip端口读取数据
func NetWorkRead(addr string) <-chan int {
	out := make(chan int)
	go func() {
		conn,err := net.Dial("tcp",addr)
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		r := ReaderSource(bufio.NewReader(conn),-1)
		for v:= range r{
			out <- v  //压入数据
		}
		close(out)
	}()
	return out
}

