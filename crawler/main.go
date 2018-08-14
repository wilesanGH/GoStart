package main

import (
	"GoStart/crawler/engine"
	"GoStart/crawler/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParseCityList,
	})

}





