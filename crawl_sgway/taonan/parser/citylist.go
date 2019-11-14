package parser

import (
	"../../engine"
	"regexp"
	"strings"
)

const cityListRe  = `<a href="(/jiaoyou/[^/]+/)">([^<]+)</a>`
//<a href="/jiaoyou/beijing/">北京交友网</a>
const w1 = "http://www.taonanw.com"
func ParserCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents,-1)  //-1代表要所有的匹配.会把上面用括号括出来的([^<]+)显示出来

	result := engine.ParseResult{}
	limit := 3
	for _, m := range matches{
		result.Items = append(result.Items,"city "+ string(m[2]))  //返回城市的名字,转成string方便(engine里面)打印
		w2 := string(m[1])
		var build strings.Builder
		build.WriteString(w1)
		build.WriteString(w2)
		w3 := build.String()

		result.Requests = append(result.Requests,engine.Request{
			Url:   w3,
			ParserFunc: ParseCity,  //再还没定义好parserfunc时先返回一个不做事的parser,这里是解析城市
		})

		limit--
		if limit ==0 {
			break
		}
	}
	return result
}
