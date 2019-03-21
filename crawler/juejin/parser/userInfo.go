package parser

//使用正则表达式提取网页上的数据
//最初的爬虫方案

//
//var zanRe = regexp.MustCompile(`获得点赞<span class="count" data-v-f6fdafdc data-v-73b88f88>([\d]+,*[\d]*)</span>次`)
//var nameRe = regexp.MustCompile(`<meta data-vue-meta="true" data-vmid="keywords" name="keywords" content="([^"]+)"/>`)
//var descriptionRe = regexp.MustCompile(`<meta data-vue-meta="true" data-vmid="description" name="description" content="([^"]*)"/>`)
//var followersRe = regexp.MustCompile(`data-v-73b88f88>关注者</div><div class="item-count" data-v-73b88f88>([\d]+,*[\d]*)</div>`)
//var followCountRe = regexp.MustCompile(`data-v-73b88f88>关注了</div><div class="item-count" data-v-73b88f88>([\d]+,*[\d]*)</div>`)
//var githubRe = regexp.MustCompile(`<a title="GitHub" href="([^"]*)" rel="nofollow noopener noreferrer" target="_blank" class="github-link link data-v-1dde3b82>" `)
//var personSiteRe = regexp.MustCompile(`<a title="个人主页" href="([^"]*)" rel="nofollow noopener noreferrer" target="_blank" class="site-link link" data-v-1dde3b82>`)
//var weiboRe = regexp.MustCompile(`<a title="微博" href="([^"]*)" rel="nofollow noopener noreferrer" target="_blank" class="weibo-link link" data-v-1dde3b82>`)
//
//
//func ParseUserHTML(contents []byte) engine.ParserResult{
//
//	profile := model.Profile{}
//	profile.Name = extractString(contents, nameRe)
//	profile.ZanCount = extractInt(contents, zanRe)
//	profile.Description = extractString(contents, descriptionRe)
//	profile.Followers = extractInt(contents, followersRe)
//	profile.FollowCount = extractInt(contents, followCountRe)
//	profile.Github = extractString(contents, githubRe)
//	profile.Weibo = extractString(contents, weiboRe)
//	profile.PersonSite = extractString(contents, personSiteRe)
//
//	result := engine.ParserResult{
//		Items: []interface{}{profile},
//	}
//	return result
//}
//
//func extractString(contents []byte, re *regexp.Regexp) string{
//	match := re.FindSubmatch(contents)
//	if len(match) >=2{
//		return string(match[1])
//	} else{
//		return ""
//	}
//}
//
//func extractInt(contents []byte, re *regexp.Regexp) int{
//	match := re.FindSubmatch(contents)
//	if len(match) >=2{
//		temp := string(match[1])
//		//将数字中的","替换
//		temp = strings.Replace(temp, ",", "", -1)
//		result, err := strconv.Atoi(temp)
//		if err == nil{
//			return result
//		}
//	}
//	return 0
//}

