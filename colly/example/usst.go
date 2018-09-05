package example

import (
	"GoStart/colly/util"
	"fmt"
	"github.com/gocolly/colly"
)

func Usst() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.usst.edu.cn"),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		if util.CheckUrlRe(link) {
			go c.Visit(e.Request.AbsoluteURL(link))
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Request URL：", r.Request.URL, "failed with response:", r, "\nError:", e)
	})

	c.Visit("http://www.usst.edu.cn")
}
