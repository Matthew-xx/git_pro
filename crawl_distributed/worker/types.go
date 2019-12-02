package worker

import (
	"../../crawl_double/engine"
	"../../crawl_double/zhenai/parser"
	"../config"
	"fmt"
	"github.com/pkg/errors"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

//不能直接套用c_double里面的types/request，因为其里面是一个parserfunc，不能直接传输到网页。
type Request struct {
	Url string
	Parser SerializedParser
}

type ParserResult struct {
	Items []engine.Item
	Requests []Request
}

//上面定义的request和engine里面的request进行转换
func SerializeRequest(r engine.Request) Request {
	name,args := r.Parser.Serialize()
	return Request{
		Url:r.Url,
		Parser:SerializedParser{
			Name:name,
			Args:args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParserResult {
	result := ParserResult{
		Items: r.Items,
	}
	for _, req := range r.Requests{
		result.Requests = append(result.Requests,SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request,error) {
	parser,err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{},err
	}
	return engine.Request{
		Url:r.Url,
		Parser:parser,
	},nil
}

func DeserializreResult(r ParserResult) engine.ParseResult {
	result := engine.ParseResult{
		Items:r.Items,
	}

	for _,req := range r.Requests{
		engineRuest,err := DeserializeRequest(req)
		if err!= nil {
			log.Printf("error deseralizing")
			continue
		}
		result.Requests = append(result.Requests,engineRuest)
	}
	return result
}

func deserializeParser(p SerializedParser) (engine.Parser,error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParserCityList,config.ParseCityList),nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity,config.ParseCity),nil
	case config.NilParser:
		return engine.NilParser{},nil
	case config.ParseProfile:
		if userName,ok := p.Args.(string);ok {
			return parser.NewProfileParser(userName),nil
		}else {return nil,fmt.Errorf("invalid args:%v"+"arg:%v",p.Args)}

	default:
		return nil,errors.New("unknow parser name")

	}
}





