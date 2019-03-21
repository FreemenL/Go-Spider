package engine

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"log"
	"strconv"
	"爬虫/crawler/fetcher"
	"爬虫/crawler/model"
)

const filename = "result.xlsx"

func Run(seeds ...Request){
	count := 0
	CreateFile(filename)
	var requests []Request
	for _, r := range seeds{
		requests = append(requests, r)
	}
	for len(requests) >0 {
		r := requests[0]
		log.Printf("fetching %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		//当获取数据失败，将请求放入队末重试
		if err!=nil{
			log.Println("Fetcher error fetching %s: %v", r.Url, err)
			requests = append(requests, requests[0])
			requests = requests[1:]
			continue
		}
		//当解析数据失败，将请求放入队末重试
		parserResult, err:= r.ParserFunc(body)
		if err!=nil{
			log.Printf("Parser error %v, retry\n", err)
			requests = append(requests, requests[0])
			requests = requests[1:]
			continue
		}
		//处理完一个后在队列里删除
		requests = requests[1:]
		//将新获取的request加入队列
		requests = append(requests, parserResult.Requests...)

		//将接口进行类型转换
		for _, item:= range parserResult.Items{
			if i, ok := item.(model.User);ok{
				log.Println(i.Username)
				WriteToFile(i)
			}
			count++
		}

		log.Printf("总数为：%d \n", count)
	}
}


//创建excel文件保存结果
func CreateFile(filename string){
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("UserInfo")
	if err!= nil{
		fmt.Println(err.Error())
	}
	row := sheet.AddRow()

	nameCell := row.AddCell()
	nameCell.Value = "用户名"

	cell := row.AddCell()
	cell.Value = "公司"

	cell = row.AddCell()
	cell.Value = "职位"

	cell = row.AddCell()
	cell.Value = "关注者"

	cell = row.AddCell()
	cell.Value = "关注数"

	cell = row.AddCell()
	cell.Value = "专栏数"

	cell = row.AddCell()
	cell.Value = "分享数（文章数）"

	cell = row.AddCell()
	cell.Value = "获赞数"

	err = file.Save(filename)
}

//将结果写入excel表中
func WriteToFile(u model.User){
	file ,err := xlsx.OpenFile(filename)

	if err!= nil{
		fmt.Errorf("open file error")
		return
	}
	sheet := file.Sheet["UserInfo"]
	row := sheet.AddRow()

	cell := row.AddCell()
	cell.Value = u.Username

	cell = row.AddCell()
	cell.Value = u.Company

	cell = row.AddCell()
	cell.Value = u.JobTitle

	cell = row.AddCell()
	count := strconv.Itoa(u.FollowersCount)
	cell.Value = count

	cell = row.AddCell()
	count = strconv.Itoa(u.FolloweesCount)
	cell.Value = count

	cell = row.AddCell()
	count = strconv.Itoa(u.PostedPostsCount)
	cell.Value = count

	cell = row.AddCell()
	count = strconv.Itoa(u.PostedEntriesCount)
	cell.Value = count

	cell = row.AddCell()
	count = strconv.Itoa(u.TotalCollectionsCount)
	cell.Value = count

	err = file.Save(filename)
	if err!=nil{
		fmt.Errorf("save error : %s", err)
	}
}