//不同的业务有不同的parser
//构建parser要根据不同的业务划分，避免混杂
//相同部分如fetcher可以放在一起
package parser

import (
	"encoding/json"
	"fmt"
	"爬虫/crawler/engine"
)

//用于解析JSON数据的结构体
type NodeMess struct {
	S int    `json:"s"`
	M string `json:"m"`
	D D      `json:"d"`
}
type Tags struct {
	ID               string    `json:"id"`
	Title            string    `json:"title"`
	EntryCount       int       `json:"entryCount"`
	SubscribersCount int       `json:"subscribersCount"`
}
type D struct {
	Tags  []Tags `json:"tags"`
	Total int    `json:"total"`
}

//调用接口，解析Json数据
func ParseNodeList(contents []byte) (engine.ParserResult, error){
	result := engine.ParserResult{}

	//获取JSON中的标签名、关注数、文章数、标签ID
	var n NodeMess
	err := json.Unmarshal(contents, &n)
	if err!=nil{
		fmt.Printf("JSON error %v\n", err)
		return engine.ParserResult{}, err
		//fmt.Printf("%s\n", contents)
	}

	tags := n.D.Tags
	tagId := ""
	pageSize := 100
	page := 0
	sort := "hotIndex"
	baseUrl := `https://timeline-merger-ms.juejin.im/v1/get_tag_entry?src=web`
	//max_page := total/pageSize
	//total := 0

	//先向第一页发送请求，获得每个标签下的用户总数，从而计算总的页数
	//拼接每个标签下第一页获取用户信息列表的URL：https://timeline-merger-ms.juejin.im/v1/
	// get_tag_entry?src=web&tagId=5597a05ae4b08a686ce56f6f&page=0&pageSize=100&sort=hotIndex
	for _, v := range tags{
		tagId = v.ID
		url := fmt.Sprintf("%s&tagId=%s&page=%d&pageSize=%d&sort=%s",baseUrl,tagId,page,pageSize,sort)
		result.Requests = append(result.Requests, engine.Request{
			Url:url,
			//通过构造一个匿名函数，将tagId传给ParseTagInit
			ParserFunc: func(c []byte) (result engine.ParserResult, e error) {
				return ParseTagInit(c, tagId)
			},
		})
	}

	//------------------------------------------------
	//拼接每个标签页的URL：https://juejin.im/tag/GO
	//url := "https://juejin.im/tag/"
	//tags := n.D.Tags
	//for _,v := range tags{
	//	result.Items = append(result.Items, v)
	//	tag_url := url+ string(v.Title)
	//	result.Requests = append(result.Requests, engine.Request{
	//		Url: tag_url,
	//		ParserFunc: engine.NilParse,
	//	})
	//}
	return result, nil
}

//不调用接口，直接从网页中用提取信息
//const nodeListRe = `<div class="title" data-v-3fabcb68>([^<]+)</div>.*?"meta subscribe"[^>]+>(\d+[^<]+)</div>.*?"meta article"[^>]+>(\d+[^<]+)</div>`
//抓取结果 title
//func ParseNodeList(contents []byte) engine.ParserResult{
//	//rx, err:= regexp.Compile(nodeListRe)
//	//if err!=nil{
//	//	panic(err)
//	//}
//
//	//使用MustCompile避免处理err，因为该正则表达式验证过正确，不会出错
//	rx := regexp.MustCompile(nodeListRe)
//	result := rx.FindAllSubmatch(contents, -1)
//	//result := rx.FindAll(contents, -1)
//	for _,v := range result{
//		for i, m:=range v{
//			if i==0{
//				continue
//			}
//			fmt.Printf("%s ", m)
//		}
//
//		fmt.Println()
//	}
//}





