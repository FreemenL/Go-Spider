//input: url string
//output: []byte, error
package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error){
	client:= &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Accept","*/*")
	request.Header.Set("Accept-Encoding","gzip, deflate, br")
	request.Header.Set("Accept-Language","zh-CN,zh;q=0.9")
	request.Header.Set("Connection","keep-alive")
	request.Header.Set("Host", "gold-tag-ms.juejin.im")
	request.Header.Set("Origin","https://juejin.im")
	request.Header.Set("Referer", "https://juejin.im/subscribe/all")
	request.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36")
	request.Header.Set("X-Juejin-Client","")
	request.Header.Set("X-Juejin-Src","web")
	request.Header.Set("X-Juejin-Token", "")
	request.Header.Set("X-Juejin-Uid","")
	resp, err := client.Do(request)
	if err!=nil{
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error status = %d", resp.StatusCode)
	}
	//将GBK转成utf8
	//通用性不好，如果不是GBK就不行
	//utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())

	//自动检测编码  x/net/html
	//将编码转为UTF-8
	buf := bufio.NewReader(resp.Body)
	e := DetermineEncoding(buf)
	utf8Reader := transform.NewReader(buf, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err!=nil{
		return nil, fmt.Errorf("Read error\n")
	}

	return all, nil
}

//利用前1024个字节判断HTML的编码类型
func DetermineEncoding(r *bufio.Reader) encoding.Encoding{
	//使用peek防止丢失前1024个bytes
	bytes, err := r.Peek(1024)
	if err!=nil{
		//若字数不够1024，返回一个默认的encoding--utf8
		//log.Fatal("fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ :=charset.DetermineEncoding(bytes, "")
	return e
}

