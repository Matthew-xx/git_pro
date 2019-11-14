package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main()  {
	resp,err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("error:status code",resp.StatusCode)
		return
	}

	//e := determineEncoding(resp.Body)
	//utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())  //自动识别及转码

	//utf8Reader := transform.NewReader(resp.Body,simplifiedchinese.GBK.NewDecoder()) //手动将utf8转码成GBK
	all,err := ioutil.ReadAll(resp.Body)  //ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n\r",all)
	printCityList(all)
}

//自动识别网页编码（utf8、GBK...
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes ,err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e,_,_ := charset.DetermineEncoding(bytes,"")
	return e
}

//正则表达式匹配
func printCityList(contents []byte)  {
	//整体打印
	//<a target="_blank" href="http://www.zhenai.com/zhenghun/shanghai" data-v-5fa74e39>上海</a>
	//re := regexp.MustCompile(`<a target="_blank" href="http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+"[^>]*>[^<]+</a>`)

	//<a data-v-5e16505f="" href="http://www.zhenai.com/zhenghun/yaan">雅安</a>
	//re := regexp.MustCompile(`<a data-v-5e16505f="" href="http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+"[^>]*>[^<]+</a>`)
	//  ``里面都是非转义字符，[^>]* 表除了>的多个其他字符，
	//matches := re.FindAll(contents,-1)  //-1代表要所有的匹配

	//子打印
	re := regexp.MustCompile(`<a target="_blank" href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents,-1)  //-1代表要所有的匹配.会把上面用括号括出来的([^<]+)显示出来
	for _, m := range matches{
		fmt.Printf("city: %s, url: %s\n",m[2],m[1]) //m[0]是整个匹配串
	}
	fmt.Printf("matches found: %d\n",len(matches))
}

//每个部分有每个部分的解析器（城市列表，城市，用户）

