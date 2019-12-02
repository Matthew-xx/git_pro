package main

import (
	"../../../crawl_double/engine"
	"../../../crawl_double/model"
	"../../rpcsupport"
	"testing"
	"time"
	"../../config"
)

func TestItemServer(t *testing.T)  {
	const host  = ":1234"
	//start Itemserver
	go serveRpc(host,"test1")  //server还没连上（listen），client就已经尝试连接，此时会出错

	time.Sleep(time.Second*5)  //应该建立通信，开启service后通知client连接

	//start Itemclient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	//call save
	item:= engine.Item{
		Url: "http://zhenai.com/u/105558",
		Type:"zhenai",
		Id:  "105558",
		Payload:model.Profile{
			Name: "小雪",
			Age: "24",
			Height: "165",
			Weight: "31",
			Marriage: "已婚",
			Income: "5-8千",
			Education: "本科",
		},
	}

	result := ""
	err = client.Call(config.ItemSaverRpc,item,&result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err :%s",result,err)
	}
}
