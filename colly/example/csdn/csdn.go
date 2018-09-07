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
	count := 0

	fName := "d:csdn_v1.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic %s\n", err)
		}
	}()

	file.WriteString("\xEF\xBB\xBF")
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"编号", "标题", "url", "关键字", "阅读数", "评论数", "发表时间", "内容"})

	//subColly
	subColly := colly.NewCollector(
		colly.AllowedDomains("blog.csdn.net"),
		//colly.Async(true),
	)
	//subColly.Limit(&colly.LimitRule{DomainGlob:"*",Parallelism:100})


	subColly.OnRequest(func(r *colly.Request) {

		//fmt.Println("subColly Visiting", r.URL.String())
	})
	subColly.OnHTML("div.blog-content-box", func(e *colly.HTMLElement) {

		csdnBlog := model.CSDN_BLOG{}
		csdnBlog.CsdnBase.Url = e.Request.URL.String()
		csdnBlog.Title = e.ChildText("h1.title-article")
		csdnBlog.Date = e.ChildText("span.time")
		csdnBlog.CsdnBase.Body = strings.Replace(strings.Replace(e.Text, "\n", "", -1), "\t", "", -1)
		csdnBlog.ReadCount = strings.Split(e.ChildText("span.read-count"), "：")[1]
		csdnBlog.Keywords = e.ChildAttrs("span.tags-box a", "data-track-view")
		//var keywordstring string
		e.ForEach("span.tags-box", func(_ int, el *colly.HTMLElement) {
			csdnBlog.Keywords = util.DeleteSpaceNTabForSlice(strings.Split(el.ChildText("a"), "\t"))
			//keywordstring = util.DeleteSpaceNTab(el.ChildText("a"))
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
			//csdnBlog.CsdnBase.Body,
		})

		count++
		log.Printf("colly count: %d",count)

		if count%100==0{
			writer.Flush()

		}
		/*log.Printf("count:%d\n", count)
		log.Println(csdnBlog)*/

	})

	subColly.OnError(func(r *colly.Response, e error) {
		fmt.Println("subColly:Request URL：", r.Request.URL, "failed with response:", r, "\nError:", e)
	})

/********************************************************************************************************/
	//parentColly
	parentColly := colly.NewCollector(
		colly.AllowedDomains("www.csdn.net",
			"blog.csdn.net"),
		//colly.Async(true),
		colly.DisallowedURLFilters(
			//regexp.MustCompile("https://blog.csdn.net/column.+"),
			regexp.MustCompile("https://blog.csdn.net/rss.h.+"),
			//regexp.MustCompile("https://blog.csdn.net/code/.+"),
		),

	)
	//parentColly.Limit(&colly.LimitRule{DomainGlob:"*",Parallelism:2})

	parentColly.OnRequest(func(r *colly.Request) {
		//fmt.Println("Visiting", r.URL.String())
	})

	parentColly.OnResponse(func(r *colly.Response) {
		reqUrl := r.Request.URL.String()
		if UrlRe.FindAllStringIndex(reqUrl, 1) == nil {
			return
		}

		//fmt.Printf("Correct: %s",reqUrl+"\n")
		if util.CheckSubUrlReapt(reqUrl) {
			go subColly.Visit(reqUrl)
		}
	})


	parentColly.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		if util.CheckUrlReapt(link) {
			func() {
				parentColly.Visit(e.Request.AbsoluteURL(link))
			}()
		}
	})

	parentColly.OnError(func(r *colly.Response, e error) {
		fmt.Println("parent:Request URL：", r.Request.URL, "failed with response:", r, "\nError:", e)
	})



	parentColly.Visit("http://www.csdn.net")
}


func GetCSDNBlog2() {
	count := 0

	fName := "d:csdn_v2.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic %s\n", err)
		}
	}()

	file.WriteString("\xEF\xBB\xBF")
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"编号", "标题", "url", "关键字", "阅读数", "评论数", "发表时间", "内容"})

	/********************************************************************************************************/
	//parentColly
	parentColly := colly.NewCollector(
		colly.AllowedDomains("www.csdn.net",
			"blog.csdn.net"),
		//colly.Async(true),
		colly.DisallowedURLFilters(
			//regexp.MustCompile("https://blog.csdn.net/column.+"),
			regexp.MustCompile("https://blog.csdn.net/rss.h.+"),
			//regexp.MustCompile("https://blog.csdn.net/code/.+"),
		),


	)
	//parentColly.Limit(&colly.LimitRule{DomainGlob:"*",Parallelism:100})

	parentColly.OnRequest(func(r *colly.Request) {
		//fmt.Println("Visiting", r.URL.String())
	})

	parentColly.OnResponse(func(r *colly.Response) {
		//fmt.Println("Return:",r.Request.URL.String())
	})

	parentColly.OnHTML("div.blog-content-box", func(e *colly.HTMLElement) {

		reqUrl := e.Request.URL.String()
		if UrlRe.FindAllStringIndex(reqUrl, 1) == nil {
			return
		}


		csdnBlog := model.CSDN_BLOG{}
		csdnBlog.CsdnBase.Url = e.Request.URL.String()
		csdnBlog.Title = e.ChildText("h1.title-article")
		csdnBlog.Date = e.ChildText("span.time")
		csdnBlog.CsdnBase.Body = strings.Replace(strings.Replace(e.Text, "\n", "", -1), "\t", "", -1)
		csdnBlog.ReadCount = strings.Split(e.ChildText("span.read-count"), "：")[1]
		csdnBlog.Keywords = e.ChildAttrs("span.tags-box a", "data-track-view")
		//var keywordstring string
		e.ForEach("span.tags-box", func(_ int, el *colly.HTMLElement) {
			csdnBlog.Keywords = util.DeleteSpaceNTabForSlice(strings.Split(el.ChildText("a"), "\t"))
			//keywordstring = util.DeleteSpaceNTab(el.ChildText("a"))
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
			//csdnBlog.CsdnBase.Body,
		})

		count++
		log.Printf("colly count: %d",count)


			writer.Flush()

		//log.Printf("count:%d\n", count)
		//log.Println(csdnBlog)

	})

	parentColly.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		if util.CheckUrlReapt(link) {
			func() {
				parentColly.Visit(e.Request.AbsoluteURL(link))
			}()
		}
	})




	parentColly.OnError(func(r *colly.Response, e error) {
		fmt.Println("parent:Request URL：", r.Request.URL, "failed with response:", r, "\nError:", e)
	})



	parentColly.Visit("http://www.csdn.net")
}

func GetCSDNBlog3() {
	c := colly.NewCollector(

	)


}