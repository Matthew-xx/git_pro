package engine

type Request struct {
	Url string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	//Items    []interface{}
	Items    []Item     //因要存放URL而改变的存储结构
}

//添加Url和ID
type Item struct {
	Url  string
	Type  string             //对应itemsaver里面存储数据的位置（数据库、表、id
	Id   string
	Payload  interface{}     //相当于存放profile里面的内容
}
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
