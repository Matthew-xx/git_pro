package model

type SearchResult struct {
	Hits int64       //总共多少
	Start int      //开始位置
	Query  string  //搜索关键字
	PrevFrom int   //上一页
	NextFrom int  //下一页
	Items []interface{}   //具体项目
}
