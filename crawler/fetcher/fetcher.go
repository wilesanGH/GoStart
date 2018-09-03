package fetcher

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/unicode"
	"log"
	"os"
	"time"
)



var rateLimiter = time.Tick(100 * time.Microsecond)

func Fetch(myUrl string,needVPN bool) ([]byte,error){
	<-rateLimiter
	//resp,err := http.Get(myUrl)
	/*urli := url.URL{}
	urlproxy, _ := urli.Parse("https://127.0.0.1:1080")
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(urlproxy),
		},
	}*/
	/*******************需要加VPN用*************************/
	if needVPN {
		os.Setenv("HTTP_PROXY", "http://127.0.0.1:1080")
		os.Setenv("HTTPS_PROXY", "https://127.0.0.1:1080")
	}
	/************************************************/
	client := &http.Client{}
	req, err := http.NewRequest("GET", myUrl,nil)
	if err != nil{
		return nil,err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; …) Gecko/20100101 Firefox/61.0")
	resp, err := client.Do(req)
	if err != nil{
		return nil,err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		return nil,fmt.Errorf("wrong status code: %d",resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	uft8Reader := transform.NewReader(bodyReader,e.NewDecoder())

	return ioutil.ReadAll(uft8Reader)

}

func determineEncoding(r *bufio.Reader) encoding.Encoding{
	bytes,err := r.Peek(1024)
	if err != nil{
		log.Printf("Fetcher error:%v",err)
		return unicode.UTF8
	}
	e,_,_ := charset.DetermineEncoding(bytes,"")
	return e
}