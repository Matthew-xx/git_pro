package baidu

import (
	"net/http"
	"net/http/httputil"
	"time"
)

//当baiduretriever里面实现了retriever中的get方法时就可以说实现了retriever这个接口，而无需再定义说明retriever接口
type Retriever struct {
	//用它实现一个接口
	UserAgent string
	TimeOut time.Duration
}

//func (r *Retriever) Get(url string) string {  //指针
func (r Retriever) Get(url string) string {
	resp,err := http.Get(url)
	if err != nil{
		panic(err)
	}

	result,err := httputil.DumpResponse(resp,true)

	resp.Body.Close()

	if err != nil {
		panic(err)
	}

	return string(result)
}

func main() {
	
}
