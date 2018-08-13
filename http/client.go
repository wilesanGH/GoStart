package main

import (
	"net/http"
	"net/http/httputil"
	"fmt"
)

func main() {
	request,err:=http.NewRequest(
		http.MethodGet,
		"https://www.baidu.com",
		nil)
	request.Header.Add("User-Agent",
		"Mozilla/5.0")
	if err!=nil{
		panic(nil)
	}


	resp,err := http.DefaultClient.Do(request)
	if err!=nil{
		panic(err)
	}

	defer
		resp.Body.Close()
	s, err := httputil.DumpResponse(resp,true)
	if err !=nil{
		panic(err)
	}

	fmt.Printf("%s\n",s)
	
}
