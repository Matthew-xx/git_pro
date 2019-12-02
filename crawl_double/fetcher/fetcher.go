package fetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"../../crawl_distributed/config"
)

var ratelimit = time.Tick(time.Second / config.Qps)
func Fetch(url string) ([]byte,error) {
	<- ratelimit   //避免访问网站速度过快遭反爬
	log.Printf("fetching url %s",url)

	//仿浏览器
	client := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	/*  //简单
	resp,err := http.Get(url)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	*/
	if resp.StatusCode != http.StatusOK {

		return nil,fmt.Errorf("wrong statuscode :%d ",resp.StatusCode)
	}

	//bodyreader := bufio.NewReader(resp.Body)
	//e := determineEncoding(bodyreader
	//utf8Reader := transform.NewReader(bodyreader ,e.NewDecoder())  //自动识别及转码
	return ioutil.ReadAll(resp.Body)  //ioutil.ReadAll(utf8Reader)
}

/*
func determineEncoding(r *bufio.Reader) encoding.Encoding { //改成指针后便不是从1025字节开始了
	bytes ,err := r.Peek(1024)  //peek是针对bufio的，对r读1024个字节然后存起来，然后上面的utf8reader则是从1025字节开始读
	if err != nil {
		return unicode.UTF8
	}
	e,_,_ := charset.DetermineEncoding(bytes,"")
	return e
}*/