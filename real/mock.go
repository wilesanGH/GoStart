package real

import (
	"net/http"
	"net/http/httputil"
	"time"
	"fmt"
)

type MockRetriever struct {
	UserAgent string
	Timeout   time.Duration
}

func (mr *MockRetriever) String() string  {
	return fmt.Sprintf("Retriever: {Contents=%s}",mr.UserAgent)
}


func (r MockRetriever) Get (url string) string  {
	resp, err := http.Get(url)
	if err!=nil{
		panic(err)
	}

	result, err:= httputil.DumpResponse(
		resp,true)
	resp.Body.Close()
	return string(result)
}

