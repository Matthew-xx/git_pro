package main

import (
	"./engine"
	"./zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		//http://www.taonanw.com
		ParserFunc: parser.ParserCityList,
	})
}