package model

import "encoding/json"

type Profile struct {
	Name       string
	Marriage   string
	Age        string
	Height     string
	Weight     string
	Income     string
	Education  string
	/*
	Occupation string  //职业
	Hokou      string  //户口
	Gender     string
	House      string
	Car        string*/
}

//解决payload中的interface存储时存成map形式而不是我们要的profile结构
//先unmarshal转换出来的东西转换成一个string，然后再将这个string转换unmarshal

func FromJsonObj(o interface{}) (Profile,error) {
	var profile Profile
	s,err := json.Marshal(o)
	if err != nil{
		return profile,err
	}

	err = json.Unmarshal(s,&profile)
	return profile,err
}

