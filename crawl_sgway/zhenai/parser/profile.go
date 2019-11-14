package parser

import (
	"../../engine"
	"../../model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([0-9]+)岁</div>`)
//var marriageRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([^<]+)</div>`)
/*
func ParseProfile(contents []byte) engine.ParseResult{
	profile := model.Profile{}
*/
func ParseProfile(contents []byte,name string) engine.ParseResult{
	profile := model.Profile{}
	profile.Name = name   //比上面注释的加了一个name，这样便可以在爬取城市时候读到的个人姓名拿过来而不用再写代码去爬
	age ,err := strconv.Atoi(extractString(contents,ageRe))
	if err != nil{
		profile.Age = age
	}

	//profile.Marriage = extractString(contents,marriageRe)

	result :=engine.ParseResult{
		Items: []interface{}{profile},
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


















