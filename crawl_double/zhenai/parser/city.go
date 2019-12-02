package parser

import (
	"../../engine"
	"regexp"
)

var (profileRe  = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
//<a href="http://album.zhenai.com/u/1139641477" target="_blank">迷迷糊糊</a>   //两个括号里面一个是链接一个是名字
	cityUrlRe= regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)
func ParseCity(contents []byte,_ string) engine.ParseResult  {
	//re := regexp.MustCompile(cityRe)
	matches := profileRe.FindAllSubmatch(contents,-1)  //-1代表要所有的匹配.会把上面用括号括出来的([^<]+)显示出来

	result := engine.ParseResult{}
	for _, m := range matches{
		url := string(m[1])
		name := string(m[2])
		//result.Items = append(result.Items,"user"+ name)  //返回城市的名字,转成string方便(engine里面)打印
		result.Requests = append(result.Requests,engine.Request{
			Url:   url,
			//ParserFunc: ParseProfile,  //再还没定义好parserfunc时先返回一个不做事的parser，这里往下解析个人
			//ParserFunc: func(contents []byte) engine.ParseResult {
				//return ParseProfile(contents,url,name)
				//对比上面的，用闭包返回可以不用修改profile里面的struct（其他地方用不上),多了一个读取用户姓名的参数
			//},
			Parser:NewProfileParser(name),
		})
	}

	//读取下一页功能
	matches_next := cityUrlRe.FindAllSubmatch(contents,-1)

	for _,m := range matches_next{
		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}

	return result
}


