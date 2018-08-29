package cl

import (
	"GoStart/crawler/engine"
	"regexp"
	"fmt"
	"GoStart/crawler/model"
	"GoStart/crawler/orm"
)

var jszqListTitleRe = regexp.MustCompile(`<h3><a href="(htm_data/[0-9]/[0-9]+/)([0-9]+).html" target="_blank" id="">([^>]*)</a></h3>  `)

var jszqListAuthorRe =regexp.MustCompile(`tid=([0-9]+)&page=e&fpage=1#a" class="f10"> ([^>]*) </a><br />by: ([^>]*)</td>`)

func ParseCLJS(content []byte) engine.ParseResult{

	matches := jszqListTitleRe.FindAllSubmatch(content,-1)
	result := engine.ParseResult{}
	cljss := model.Cljss{}
	for _,m := range matches{
		profile := model.Cljs{}
		profile.Id = string(m[2])
		profile.Title = string(m[3])
		profile.Url = "http://www.t66y.com/"+string(m[1])+string(m[2])+".html"
		cljss.MyProfile = append(cljss.MyProfile, profile)
		fmt.Printf("id:%s,title: %s,url: %s",string(m[2]),string(m[3]),"http://www.t66y.com/"+string(m[1])+string(m[2])+".html")
		fmt.Println()
	}

	matches2 := jszqListAuthorRe.FindAllSubmatch(content,-1)
	for _,m := range matches2{
		profile := model.Cljs{}
		profile.Id = string(m[1])
		profile.PTime = string(m[2])
		profile.Author = string(m[3])
		for i,v := range cljss.MyProfile{
			if v.Id == profile.Id{
				cljss.MyProfile[i].Author = profile.Author
				cljss.MyProfile[i].PTime = profile.PTime
				continue
			}
		}
		fmt.Printf("id:%s,publicTime:%s,author:%s",string(m[1]),string(m[2]),string(m[3]))
		fmt.Println()

	}
	n,err := orm.Engine.Table("cljs").Insert(&cljss.MyProfile)
	if err != nil{
		panic(err)
	}
	fmt.Printf("插入了%d数据",n)

	return result
}


