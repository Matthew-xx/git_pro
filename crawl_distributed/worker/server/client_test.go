package main

import (
	"../../config"
	"../../rpcsupport"
	"../../worker"
	"fmt"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T)  {
	const host  = ":9000"
	go rpcsupport.ServeRpc(host,worker.CrawlService{})
	time.Sleep(time.Second)

	client,err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url:"http://www.album.zhenai.com/15412",
		Parser:worker.SerializedParser{
			Name:config.ParseProfile,
			Args:"小雪",
		},
	}

	var result worker.ParserResult
	err = client.Call(config.CrawlServiceRpc,req,&result)
	if err != nil {
		t.Error(err)
	}else {
		fmt.Println(result)
	}
}
