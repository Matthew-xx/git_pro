package scheduler

import "../engine"

type SimpleScheduler struct {
	wokerChan chan engine.Request  //每个worker共用一个channel
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request{
	return s.wokerChan
}

func (s *SimpleScheduler) Run() {
	s.wokerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}

/*
func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.wokerChan = c
}*/

func (s *SimpleScheduler) Submit(r engine.Request) {
	//send request down to worker chan
	//s.wokerChan <- r  //为避免无人接收导致卡死，可以使用go
	go func() { s.wokerChan <- r }()  //开一个gorutine去submit

	//在concurrent中，将收到的result给out（最下面),在30行后面打印完再submit，而这边开的gorutine很快就返回掉，不会卡死
}




