package main

import (
	"./engine"
	"./persist"
	"./scheduler"
	"./zhenai/parser"
)

func main() {
	ItemChan,err :=persist.ItemSaver("dating_profile")  //
	if err != nil{
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:   ItemChan ,
		RequestProcessor:engine.Worker,
	}
	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		//http://www.taonanw.com
		Parser: engine.NewFuncParser(parser.ParserCityList,"ParserCityList"),
	})
}