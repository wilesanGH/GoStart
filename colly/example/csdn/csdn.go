package csdn

import (
	"GoStart/colly/example/csdn/model"
	"GoStart/colly/util"
	"encoding/csv"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
	"regexp"
	"strings"
)

var UrlRe = regexp.MustCompile(`https://blog.csdn.net/[a-zA-z0-9_]+/article/details/[0-9]+`)

var blogs []model.CSDN_BLOG

func GetCSDNBlog() {
	fName := "/opt/csdn_v1.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"编号", "标题", "url", "关键字", "阅读数", "评论数", "发表时间", "内容"})

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

		//fmt.Println("rightUrlColly Visiting", r.URL.String())
	})
	c.OnResponse(func(r *colly.Response) {
		reqUrl := r.Request.URL.String()
		if UrlRe.FindAllStringIndex(reqUrl, 1) == nil {
			return
		}

		//fmt.Printf("Correct: %s",reqUrl+"\n")
		if util.CheckSubUrlReapt(reqUrl) {
			rightUrlColly.Visit(reqUrl)
		}
	})

	count := 0

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		if util.CheckUrlReapt(link) && count < 100000 {
			c.Visit(e.Request.AbsoluteURL(link))
		}
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Request URL：", r.Request.URL, "failed with response:", r, "\nError:", e)
	})

	//writer.Write([]string{"编号", "标题", "url", "关键字", "阅读数", "评论数", "发表时间", "内容"})

	rightUrlColly.OnHTML("div.blog-content-box", func(e *colly.HTMLElement) {

		csdnBlog := model.CSDN_BLOG{}
		csdnBlog.CsdnBase.Url = e.Request.URL.String()
		csdnBlog.Title = e.ChildText("h1.title-article")
		csdnBlog.Date = e.ChildText("span.time")
		csdnBlog.CsdnBase.Body = strings.Replace(strings.Replace(e.Text, "\n", "", -1), "\t", "", -1)
		csdnBlog.ReadCount = strings.Split(e.ChildText("span.read-count"), "：")[1]
		csdnBlog.Keywords = e.ChildAttrs("span.tags-box a", "data-track-view")
		e.ForEach("span.tags-box", func(_ int, el *colly.HTMLElement) {
			csdnBlog.Keywords = util.DeleteSpaceNTabForSlice(strings.Split(el.ChildText("a"), "\t"))
		})

		csdnBlog.CsdnBase.Id = csdnBlog.CsdnBase.Url[strings.LastIndex(csdnBlog.CsdnBase.Url, "/")+1:]
		writer.Write([]string{
			csdnBlog.CsdnBase.Id,
			csdnBlog.Title,
			csdnBlog.CsdnBase.Url,
			fmt.Sprint(csdnBlog.Keywords)[:len(fmt.Sprint(csdnBlog.Keywords))-1][1:],
			csdnBlog.ReadCount,
			csdnBlog.CommentCount,
			csdnBlog.Date,
			csdnBlog.CsdnBase.Body,
		})
		count++
		log.Printf("count:%d\n", count)
		log.Println(csdnBlog)

	})

	rightUrlColly.OnError(func(r *colly.Response, e error) {
		fmt.Println("Request URL：", r.Request.URL, "failed with response:", r, "\nError:", e)
	})

	c.Visit("http://www.csdn.net")
}
