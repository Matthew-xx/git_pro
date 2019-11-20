package engine

import (
	"../model"
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	//ConfigureMasterWorkerChan(chan Request)
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)  //原本在schedule中，避免笨重，单独拿出来
}

func (e *ConcurrentEngine) Run (seeds ...Request) {

	//in:= make(chan Request)
	out:= make(chan ParseResult)
	//e.Scheduler.ConfigureMasterWorkerChan(in)
	e.Scheduler.Run()

	for i:=0;i<e.WorkerCount;i++{
		createWorker(e.Scheduler.WorkerChan(),out,e.Scheduler)
		//向e.Scheduler.WorkerChan()要workerchan
	}

	for _,r := range seeds{
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		result := <- out  //收Out
		for _,item := range result.Items{
			if _,ok := item.(model.Profile);ok {
				log.Printf("GOT item #%d: %v", itemCount,item)
				itemCount++
			}
		}
		for _,request := range result.Requests{
			e.Scheduler.Submit(request)  //再将request送给schedule调度器
		}
	}
}



func createWorker(in chan Request,out chan ParseResult,ready ReadyNotifier)  {
	//in := make(chan Request)
	go func() {
		for {
			//tell scheduler i'm ready
			ready.WorkerReady(in)
			request := <- in
			result,err := worker(request)
			if err!= nil {
				continue
			}
			out <- result  //将result送去out
		}
	}()
}



