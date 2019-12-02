package main

import (
	"../../rpcsupport"
	"../../worker"
	"flag"
	"fmt"
	"log"
)

//建命令行参数，不同的WorkerPort
//在命令行窗口（cmd)输入：go run worker.go --port=9000
var port = flag.Int("port",0,"the port for me to listen on")

func main()  {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
	}
	//log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d",config.WorkerPort0),worker.CrawlService{}))
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d",*port),worker.CrawlService{}))
}
