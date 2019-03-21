package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"爬虫/crawler/engine"
	"爬虫/crawler/model"
)

//for _, v := range user.D.Entrylist{
//	fmt.Printf("%s %s %s ，关注者：%d ，关注了：%d ，专栏数：%d ，分享数：%d ，获赞数：%d",v.User.Username, v.User.Company, v.User.JobTitle,
//		v.User.FollowersCount, v.User.FolloweesCount, v.User.PostedPostsCount, v.User.PostedEntriesCount, v.User.TotalCollectionsCount)
//	fmt.Println()
//}
type userinfo struct {
	D struct {
		Total     int `json:"total"`
		Entrylist []struct {
			model.User   `json:"user,omitempty"`
		} `json:"entrylist"`
	} `json:"d"`
}


//获取每个Tag首页的用户信息
//分离主要目的是用来获取Total总数，用于计算每个Tag用户信息的总页数
func ParseTagInit(contents []byte, tagId string) (engine.ParserResult, error){
	result := engine.ParserResult{}
	firstPageUser := userinfo{}
	err := json.Unmarshal(contents, &firstPageUser)
	if err!=nil{
		fmt.Printf("JSON error %v\n", err)
		return engine.ParserResult{}, err
	}

	//该标签下的用户总数
	total := firstPageUser.D.Total
	pageSize := 100
	page := 1
	sort := "hotIndex"
	baseUrl := `https://timeline-merger-ms.juejin.im/v1/get_tag_entry?src=web`
	//计算最大页数
	maxPage := total/pageSize

	for _,v := range firstPageUser.D.Entrylist{
		result.Items = append(result.Items, v.User)
		for page<=maxPage{
			url := fmt.Sprintf("%s&tagId=%s&page=%d&pageSize=%d&sort=%s",baseUrl,tagId,page,pageSize,sort)
			result.Requests = append(result.Requests, engine.Request{
				Url:url,
				ParserFunc:ParseUser,
			})
			page++
		}
	}

	return result, nil
}

//获取用户信息
func ParseUser(contents []byte) (engine.ParserResult, error)  {
	result := engine.ParserResult{}
	user := userinfo{}
	err := json.Unmarshal(contents, &user)
	if err!=nil{
		fmt.Printf("JSON error %v\n", err)
		return engine.ParserResult{}, err
	}

	for _, v:= range user.D.Entrylist{
		log.Println("result: = ", v.User)
		result.Items = append(result.Items, v.User)
	}
	return result, nil
}