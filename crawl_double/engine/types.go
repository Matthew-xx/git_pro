package engine


type ParserFunc func(contents []byte,url string) ParseResult
//解析器序列化和反序列化。函数不能直接传到网络上，需在worker中序列化为json，json传到engine中再反序列化为函数

type Parser interface {
	Parse(contents []byte,url string) ParseResult
	Serialize() (name string,args interface{})
}

type Request struct {
	Url string
	Parser Parser   //把Parser从之前的parserfunc变为现在的接口，这个接口有两个实现，一是FuncParser
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

//实现parser的inte
type NilParser struct {}

func (NilParser) Parse(_ []byte,_ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string,args interface{}){
	return "Nilparser",nil
}

type FuncParser struct {
	parser ParserFunc
	name string
}

func (f *FuncParser) Parse(contents []byte,url string) ParseResult{
	return f.parser(contents,url)
}

func (f *FuncParser) Serialize() (name string,args interface{}){
	return f.name,nil
}

func NewFuncParser(p ParserFunc,name string) *FuncParser{
	return &FuncParser{
		parser:p,
		name:name,
	}
}

