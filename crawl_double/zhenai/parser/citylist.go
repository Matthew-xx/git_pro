package parser

import (
	"../../engine"
	"regexp"
)

const cityListRe  = `<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]*>([^<]+)</a>`
//<a href="http://www.zhenai.com/zhenghun/akesu" data-v-5e16505f>阿克苏</a>

//const cityListRe  = `<a target="_blank" href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]*>([^<]+)</a>`
//<a data-v-5e16505f="" href="http://www.zhenai.com/zhenghun/aba">阿坝</a>
func ParserCityList(contents []byte,_ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents,-1)  //-1代表要所有的匹配.会把上面用括号括出来的([^<]+)显示出来

	result := engine.ParseResult{}
	limit := 20
	for _, m := range matches{
		//result.Items = append(result.Items,"city"+ string(m[2]))  //返回城市的名字,转成string方便(engine里面)打印
		result.Requests = append(result.Requests,engine.Request{
			Url:   string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"), //再还没定义好parserfunc时先返回一个不做事的parser,这里是解析城市
		})  //看到ParseCity"，就调用该函数
		limit--
		if limit ==0 {
			break
		}
	}
	return result
}
