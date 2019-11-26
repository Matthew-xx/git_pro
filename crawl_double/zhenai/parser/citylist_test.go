package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T)  {
	contents,err := ioutil.ReadFile("citylist_test_data.html")
	//contents,err := fetcher.Fetch("http://www.zhenai.com/zhenghun")  //先从网站读取后保存成文档以后测试（避免在线网站出错

	if err != nil{
		panic(err)
	}

	//fmt.Printf("%s\n",contents)
	result := ParserCityList(contents)
	const resultsize  =  70

	expectedUrls := []string{"","","",}
	//expectedCities := []string{"","","",}
	
	if len(result.Requests) != resultsize {
		t.Errorf("result should have %d" + "requests; but have %d",resultsize,len(result.Requests))  //对比应该有的数量和实际读取到的数量判断是否正确读取
	}

	for i,url := range expectedUrls{
		if result.Requests[i].Url != url{
			t.Errorf("expected url %d: %s;but" + "was %s",i,url,result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultsize {
		t.Errorf("result should have %d" + "requests; but have %d",resultsize,len(result.Items))  //对比应该有的数量和实际读取到的数量判断是否正确读取
	}
	/*
	for i,city := range expectedCities{
		if result.Items[i].(string) != city{
			t.Errorf("expected city %d: %s;but" + "was %s",i,city,result.Items[i].(string))
		} //数量对得上再看下级（各城市）是否对得上
	}*/
}