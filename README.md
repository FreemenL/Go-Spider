# Go

项目详细说明：https://guiguiyo.cn/2019/03/21/Go%E5%AE%9E%E6%88%98-%E7%88%AC%E5%8F%96%E6%8E%98%E9%87%91%E7%BD%91%E7%94%A8%E6%88%B7%E4%BF%A1%E6%81%AF/

Go爬虫爬取掘金用户信息

​	爬取每个掘金网每个标签下文章作者信息，结果保存在result.xlsx中

​	需要安装**xlsx**库：go get github.com/tealeg/xlsx

​	engine.go用于管理任务队列；fetch.go用于获取网站界面；crawler\juejin\parser 用于提取界面数据

​	停止程序时尽量不要在写入文件时停止，否则可能会导致文件写入失败无法读取

