package main

import (
	"./pipelineMiddleware"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

//本地，多线程，分布式

//分布式（网络版
func createNetWorkPipline(filename string,filesize int,chunkcount int) <- chan int {
	file,_ := os.Create(filename)
	defer file.Close()

	mypipe := pipelineMiddleware.RandSource(filesize/8)  //管道装随机数
	writer := bufio.NewWriter(file)  //xr
	pipelineMiddleware.WriteSlink(writer,mypipe) //将mypipe写入
	writer.Flush() //刷新

	chunksize := filesize/chunkcount  //每段大小
	sortAddr := []string{}  //地址集合
	pipelineMiddleware.Init()
	file,err := os.Open(filename)  //打开文件
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for i:=0; i<chunkcount; i++ {
		file.Seek(int64(i*chunksize),0)  //移动文件指针的位置
		source := pipelineMiddleware.ReaderSource(bufio.NewReader(file),chunksize) //读取数据
		addr := ":" + strconv.Itoa(7000+i)  //开辟地址

		pipelineMiddleware.NetWorkWrite(addr,pipelineMiddleware.InMemorySort(source))  //写入到分布式主机
		sortAddr = append(sortAddr,addr)  //地址复制，保存
	}
	
	sortresults := [] <- chan int{}
	for _,addr := range sortAddr{
		sortresults = append(sortresults,pipelineMiddleware.NetWorkRead(addr))  //读取数据
	}

	return pipelineMiddleware.MergeN(sortresults...)
}


//多线程--调用中间件完成
func createPipline(filename string,filesize int,chunkcount int) <- chan int {
	file,_ := os.Create(filename)
	defer file.Close()

	mypipe := pipelineMiddleware.RandSource(filesize/8)  //管道装随机数
	writer := bufio.NewWriter(file)  //xr
	pipelineMiddleware.WriteSlink(writer,mypipe) //将mypipe写入
	writer.Flush() //刷新

	chunksize := filesize/chunkcount  //数量
	sortresults := []<- chan int{}  //排序结果，每一个元素是一个管道
	pipelineMiddleware.Init() //初始化

	file,err := os.Open(filename)  //打开文件
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for i:=0; i<chunkcount; i++ {

		file.Seek(int64(i*chunksize),0)  //跳到文件指针
		source := pipelineMiddleware.ReaderSource(bufio.NewReader(file),chunksize) //读取
		sortresults = append(sortresults,pipelineMiddleware.InMemorySort(source)) //对结果排序
	}  //调用多个内存排序，并归并
	return pipelineMiddleware.MergeN(sortresults...)
}

//写入文件
func writeTofile(in <- chan int,filename string)  {
	file,err := os.Create(filename)  //打开文件
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipelineMiddleware.WriteSlink(writer,in)  //写入数据
}

//显示文件
func showFile(filename string)  {
	file,err := os.Open(filename)  //打开文件
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipelineMiddleware.ReaderSource(bufio.NewReader(file),-1)

	countx := 0
	for v:= range  p{
		fmt.Println(v)
		countx++
		if countx > 1000{
			break
		}
	}
}

func mainc()  {
	go func() {
		time.Sleep(time.Second*10)
	}()

	p := createPipline("big.in",800000,4) //80万数据分4段
	writeTofile(p,"big.out")
	showFile("big.out")
}

func main()  {
	go func() {
		time.Sleep(time.Second*10)
	}()

	p := createNetWorkPipline("big.in",800000,4) //80万数据分4段
	writeTofile(p,"big.out")
	showFile("big.out")
}

