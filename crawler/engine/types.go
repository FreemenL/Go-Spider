package engine

//定义类型
type Request struct {
	Url string
	ParserFunc func([]byte) (ParserResult, error)
}

type ParserResult struct {
	Requests []Request
	Items []interface{}
}

