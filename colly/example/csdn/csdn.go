package csdn

import (
	"github.com/gocolly/colly"
	"fmt"
	"GoStart/colly/util"
	"regexp"
	"GoStart/colly/example/csdn/model"
	"strings"
	"strconv"
)

var UrlRe = regexp.MustCompile(`https://blog.csdn.net/[a-zA-z0-9]+/article/details/[0-9]+`)

var blogs []model.CSDN_BLOG

func GetCSDNBlog()  {
	c := colly.NewCollector(
		colly.AllowedDomains("www.csdn.net",
			"blog.csdn.net"),
	)

	rightUrlColly := colly.NewCollector(
		colly.AllowedDomains("blog.csdn.net"),
	)
	c.OnRequest(func(r *colly.Request) {
		//fmt.Println("Visiting", r.URL.String())
	})

	rightUrlColly.OnRequest(func(r *colly.Request) {

		fmt.Println("rightUrlColly Visiting", r.URL.String())
	})
	c.OnResponse(func(r *colly.Response) {
		reqUrl := r.Request.URL.String()
		if UrlRe.FindAllStringIndex(reqUrl,1) == nil{
			return
		}

		fmt.Printf("***********************Correct: %s",reqUrl)
		rightUrlColly.Visit(reqUrl)

	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		if util.CheckUrlRe(link) {
			c.Visit(e.Request.AbsoluteURL(link))
		}
	})


	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Request URL：", r.Request.URL, "failed with response:", r, "\nError:", e)
	})

	rightUrlColly.OnHTML("div.blog-content-box", func(e *colly.HTMLElement) {
		var err error
		csdnBlog := model.CSDN_BLOG{}
		csdnBlog.CsdnBase.Url = e.Request.URL.String()
		csdnBlog.CsdnBase.Body = e.Text
		csdnBlog.Title = e.ChildText("h1.title-article")
		csdnBlog.Date = e.ChildText("span.time")
		//str := strings.Split(e.ChildText("span.read-count"),"：")[1]
		csdnBlog.ReadCount,err = strconv.Atoi(strings.Split(e.ChildText("span.read-count"),"：")[1])
		if err != nil{
			panic(err)
		}


	})

	rightUrlColly.OnError(func(r *colly.Response, e error) {
		fmt.Println("Request URL：", r.Request.URL, "failed with response:", r, "\nError:", e)
	})




	c.Visit("http://www.csdn.net")
}