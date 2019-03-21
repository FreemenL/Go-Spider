# Go

Go爬虫爬取掘金用户信息

​	爬取每个掘金网每个标签下文章作者信息，结果保存在result.xlsx中

​	需要安装**xlsx**库：go get github.com/tealeg/xlsx

​	engine.go用于管理任务队列；fetch.go用于获取网站界面；crawler\juejin\parser 用于提取界面数据

​	

