package main

import (
	"./engine"
	"./scheduler"
	"./zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 5,
	}
	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		//http://www.taonanw.com
		ParserFunc: parser.ParserCityList,
	})
}