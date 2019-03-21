package main

import (
	"爬虫/crawler/engine"
	"爬虫/crawler/juejin/parser"
)

func main(){
	//从这四个界面获取所有的Tag信息
	url1 := "https://gold-tag-ms.juejin.im/v1/tags/type/hot/page/1/pageSize/40"
	url2 := "https://gold-tag-ms.juejin.im/v1/tags/type/hot/page/2/pageSize/40"
	url3 := "https://gold-tag-ms.juejin.im/v1/tags/type/hot/page/3/pageSize/40"
	url4 := "https://gold-tag-ms.juejin.im/v1/tags/type/hot/page/4/pageSize/40"

	engine.Run(
		engine.Request{
			Url:url1,
			ParserFunc: parser.ParseNodeList,},
		engine.Request{
			Url:url2,
			ParserFunc: parser.ParseNodeList,},
		engine.Request{
			Url:url3,
			ParserFunc: parser.ParseNodeList,},
		engine.Request{
			Url:url4,
			ParserFunc: parser.ParseNodeList,},
		)


}


