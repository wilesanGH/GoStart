package engine

import (
	"GoStart/crawler/fetcher"
	"log"
)

func Run(seeds ...Request){
	var requests []Request
	for _,r := range seeds{
		requests = append(requests,r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("fetching %s",r.Url)
		body, err := fetcher.Fetch(r.Url,r.NeedVPN)
		//fmt.Printf("%s",body)
		if err != nil{
			log.Printf("Fetcher: error fetching url %s: %v",r.Url,err)
		}
		parseResult := r.ParserFunc(body)
		requests = append(requests,parseResult.Requests...)

		for _,item := range parseResult.Items {
			log.Println(item)
		}
	}
}
