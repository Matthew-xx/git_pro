package engine

import (
	"../fetcher"
	"log"
)

type SimpleEngine struct {}
func (e SimpleEngine)Run(seeds ...Request)  { //传入种子页面并放入requests列表
	var requests []Request
	for _, r := range seeds{
		requests = append(requests,r)
	}

	for len(requests)>0 {
		r := requests[0]
		requests = requests[1:]  //拿出第一个request

		parseResult ,err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests,parseResult.Requests...) //后面三个点表示将读取到的内容一个个展开加进去,不然要像下行这么麻烦
		//requests = append(requests,parseResult.Requests[0],parseResult.Requests[1],...)

		for _,item := range parseResult.Items{
			log.Printf("got item %s",item)
		}

	}
}

//Run 获取seeds，维护request队列，对每一个request去fetcher，将fetcher的结果放在body里，再对body进行parserfunc获取parseresult，最后放入requests

//并发版，将fetcher和parse提取出来合在一块（因为fetcher的输出是parser的输入）,然后开多个worker并发
func worker(r Request) (ParseResult,error)  {
	log.Printf("etching %s",r.Url)
	body ,err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("fetcher :error" + "fetching url %s: %v",r.Url,err)
		return ParseResult{},err
	}
	return r.ParserFunc(body),nil
}
