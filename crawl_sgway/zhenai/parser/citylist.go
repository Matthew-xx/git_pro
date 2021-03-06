package parser

import (
	"../../engine"
	"regexp"
)

const cityListRe  = `<a target="_blank" href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]*>([^<]+)</a>`

func ParserCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents,-1)  //-1代表要所有的匹配.会把上面用括号括出来的([^<]+)显示出来

	result := engine.ParseResult{}
	limit := 3
	for _, m := range matches{
		result.Items = append(result.Items,"city"+ string(m[2]))  //返回城市的名字,转成string方便(engine里面)打印
		result.Requests = append(result.Requests,engine.Request{
			Url:   string(m[1]),
			ParserFunc: ParseCity,  //再还没定义好parserfunc时先返回一个不做事的parser,这里是解析城市
		})
		limit--
		if limit ==0 {
			break
		}
	}
	return result
}
