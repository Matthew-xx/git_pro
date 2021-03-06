package parser

import (
	"../../engine"
	"regexp"
	"strings"
)

const cityRe  = `<a target="_blank" href="(/u_[\d]+)" title="(.*)" class="netname"> <b>([^<]+)</b> </a>`
//<a href="http://album.zhenai.com/u/1139641477" target="_blank">迷迷糊糊</a>   //两个括号里面一个是链接一个是名字
//<a target="_blank" href="/u_16500659" title="马克西" class="netname">


const s1 = "http://www.taonanw.com"
func ParseCity(contents []byte) engine.ParseResult  {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents,-1)  //-1代表要所有的匹配.会把上面用括号括出来的([^<]+)显示出来

	result := engine.ParseResult{}
	for _, m := range matches{
		name := string(m[3])
		s2 := string(m[1])
		var build strings.Builder
		build.WriteString(s1)
		build.WriteString(s2)
		s3 := build.String()

		result.Items = append(result.Items,"user "+ name)  //返回城市的名字,转成string方便(engine里面)打印
		result.Requests = append(result.Requests,engine.Request{
			Url:   s3,
			//ParserFunc: ParseProfile,  //再还没定义好parserfunc时先返回一个不做事的parser，这里往下解析个人
			ParserFunc: func(contents []byte) engine.ParseResult {
				return ParseProfile(contents,name)
				//对比上面的，用闭包返回可以不用修改profile里面的struct（其他地方用不上),多了一个读取用户姓名的参数
			},
		})
	}
	return result
}
