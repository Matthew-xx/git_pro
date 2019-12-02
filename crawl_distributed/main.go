package main

import (
	"../crawl_double/engine"
	"../crawl_double/scheduler"
	"../crawl_double/zhenai/parser"
	"./config"
	itemsaver "./persist/client"
	"./rpcsupport"
	worker "./worker/client"
	"flag"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host","","itemsaver host")
	workerHosts = flag.String("worker_hosts","","worker hosts (comma seperated)")
)
//然后在命令行输入：go run main.go --itemsaver_host=":123" --worker_hosts=":9000,:9001"  #运行两个worker服务器

func main() {
	flag.Parse()
	ItemChan,err :=itemsaver.ItemSaver(*itemSaverHost)  //其他都没变，只是itemserver调用远程rpc
	if err != nil{
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts,","))  //是切片，需解析出来(如上面运行两个端口用 ","隔开

	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:       &scheduler.QueuedScheduler{},
		WorkerCount:     100,
		ItemChan:          ItemChan ,
		RequestProcessor:  processor,  //100个任务给n个processor
	}
	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		//http://www.taonanw.com
		Parser: engine.NewFuncParser(parser.ParserCityList,config.ParseCityList),
	})
}

//建一个pool池
func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _,h := range hosts{
		client,err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients,client)
			log.Printf("connected to %s:",h)
		}else {
			log.Printf("error connecting to %s: %v",h,err)
		}
	}

	out := make(chan *rpc.Client)  //往channel里分发clients
	go func() {
		//可以随机分发也可以轮流顺序分发
		for  {
			for _,client := range clients{
				out <- client
			}
		} //嵌套for循环，一直分发
	}()

	return out
}

