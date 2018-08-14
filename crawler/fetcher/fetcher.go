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
)

func Fetch(url string) ([]byte,error){
	resp,err := http.Get(url)
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