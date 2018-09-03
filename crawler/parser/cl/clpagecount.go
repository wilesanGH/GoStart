package cl

import (
	"regexp"
	"strconv"
	"GoStart/crawler/engine"
	"time"
)

var pageCountRe = regexp.MustCompile(`style="width:50px;" value="1/([0-9]+)"`)
var pageCount int

func GetPageCount(content []byte)engine.ParseResult{
	var err error

	count := "150"//extractString(content,pageCountRe)
	pageCount,err = strconv.Atoi(count)
	if err != nil{
		panic(err)
	}
	for i:=1;i<=pageCount;i++ {
		time.Sleep(1000)
		pageId := strconv.Itoa(i)
		engine.SimpleEngine.Run(engine.Request{
			Url:"http://www.t66y.com/thread0806.php?fid=7&search=&page="+string(pageId),
			ParserFunc:ParseCLJS,
			NeedVPN:true,

		})
	}
	return engine.ParseResult{}
}

