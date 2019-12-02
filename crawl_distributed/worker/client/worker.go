package client

import (
	"../../../crawl_double/engine"
	"../../config"
	"../../worker"
	"net/rpc"
)
func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	//client,err := rpcsupport.NewClient(fmt.Sprintf(":%d",config.WorkerPort0))  //单个服务器
	//if err != nil {	return nil, err }

	return func(req engine.Request) ( engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)

		var sResult worker.ParserResult
		cClient := <- clientChan   //c从clientChan 拿一个client服务器然后调用就可以了
		err := cClient.Call(config.CrawlServiceRpc,sReq,&sResult)
		if err != nil {
			return engine.ParseResult{},err
		}
		return worker.DeserializreResult(sResult),nil
	}
}
