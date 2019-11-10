package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main()  {
	request,err := http.NewRequest(
		http.MethodGet,"http://www.imooc.com",nil)
	request.Header.Add("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	//设置头部访问手机版的网页

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("redirect:",req)
			return nil
		},
	}  //每一次的目标放req里， 所有重定向的路径都放在via里面。返回nil就让它重定向，返回错误就终止重定向
	resp,err := client.Do(request)  //使用自定义client
	//resp,err := http.DefaultClient.Do(request)  //使用默认的client
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s,err := httputil.DumpResponse(resp,true)
	if err != nil{
		panic(err)
	}
	fmt.Printf("%s\n",s)
}
