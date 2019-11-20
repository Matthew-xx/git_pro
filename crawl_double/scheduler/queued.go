package scheduler

import (
	"../engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request  //一个总的workerChan,里面是worker（chan engine.Request)
}

func (s * QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request)  {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <-w
}
/*
func (s *QueuedScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	panic("implement me")
}*/

func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request   //声明两个对列
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ)>0 && len(workerQ)>0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}

			select {//同时取，哪个有（发生）就取哪个并缓存
			case r := <- s.requestChan:
				requestQ = append(requestQ,r)
			case w := <- s.workerChan:
				workerQ = append(workerQ,w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]  //同时两个队列都有数据时，就传。传了后截掉，当队列为空时只会select前面两个case

			}
		}
	}()
}



