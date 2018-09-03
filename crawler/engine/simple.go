package engine

import (
	"GoStart/crawler/fetcher"
	"log"
)

type SimpleEngine struct {

}

func (e SimpleEngine) Run(seeds ...Request){
	var requests []Request
	for _,r := range seeds{
		requests = append(requests,r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult,err := Worker(r)
		if err != nil{
			panic(err)
		}
		requests = append(requests,parseResult.Requests...)

		for _,item := range parseResult.Items {
			log.Println(item)
		}
	}
}

func  Worker(r Request) (ParseResult,error){
	log.Printf("fetching %s",r.Url)
	body, err := fetcher.Fetch(r.Url,r.NeedVPN)
	//fmt.Printf("%s",body)
	if err != nil{
		log.Printf("Fetcher: error fetching url %s: %v",r.Url,err)
		return ParseResult{},err
	}
	return r.ParserFunc(body),nil
}
