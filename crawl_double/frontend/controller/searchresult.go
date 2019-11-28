package controller

import (
	"../view"
	"../model"
	"context"
	"github.com/olivere/elastic"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"../../engine"
)

type SearchResultHandler struct {
	view view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client,err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	return SearchResultHandler{
		view:view.CreateSearchResultView(template),
		client:client,
	}
}

//实现 localhost:8888/search?q=未婚&from=10     #实现查询且有翻页功能，每页10个
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter,req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))  //去空格
	from ,err := strconv.Atoi(req.FormValue("from"))

	if err != nil {
		from = 0 //忽略错误
	}
	//fmt.Fprintf(w,"q=%s,from=%d",q,from)

	//var page model.SearchResult
	page,err := h.getSearchResult(q,from)
	if err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
	}

	err = h.view.Render(w,page)
	if err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
	}
}

func (h SearchResultHandler) getSearchResult(q string,from int) (model.SearchResult,error){
	var result model.SearchResult
	result.Query = q    //将搜索词填入搜索框（配合HTML
	resp,err := h.client.Search("dating_profile").
		Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).From(from).Do(context.Background())

	if err != nil {
		return result,err
	}

	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)

	return result,nil
}

//搜索字符串重写
func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q,"Payload.$1:")  //$1表上面([A-Z][a-z]*).....替换搜索框里面的内容，搜索时不用将字段名输入
}


