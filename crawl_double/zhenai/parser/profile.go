package parser

import (
	"../../engine"
	"../../model"
	"regexp"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([0-9]+)岁</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>月收入:([^<]+)</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([0-9]+)cm</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([0-9]+)kg</div>`)
var educationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+科)</div>`)

//获取id
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)
//<div class="m-btn purple" data-v-8b1eac0c>38岁</div>
//<div class="m-btn purple" data-v-8b1eac0c>离异</div>
//<div class="m-btn purple" data-v-8b1eac0c>37岁</div>
//<div class="m-btn purple" data-v-8b1eac0c>中专</div>
/*
func ParseProfile(contents []byte) engine.ParseResult{
	profile := model.Profile{}
*/
func ParseProfile(contents []byte,url string,name string) engine.ParseResult{
	profile := model.Profile{}
	profile.Name = name   //比上面注释的加了一个name，这样便可以在爬取城市时候读到的个人姓名拿过来而不用再写代码去爬
	/*
	age ,err := strconv.Atoi(extractString(contents,ageRe))
	if err != nil{
		profile.Age = age
	}*/

	profile.Marriage = extractString(contents,marriageRe)
	profile.Age = extractString(contents,ageRe)
	profile.Height = extractString(contents,heightRe)
	profile.Weight = extractString(contents,weightRe)
	profile.Income = extractString(contents,incomeRe)
	profile.Education = extractString(contents,educationRe)

	result :=engine.ParseResult{
		Items: []engine.Item{
			{
				Url:url,
				Type:"zhenai",
				Id: extractString([]byte(url),idUrlRe),
				Payload:profile,
			},
		},
	}
	return result
}

func extractString(contents []byte, r *regexp.Regexp) string {
	match := r.FindSubmatch(contents)

	if len(match) >= 2{  //因为要取第一个（match[0}是全的，match[1]是括号里面的，所以最少2
		return string(match[1])
	}else {
		return ""
	}
}


















