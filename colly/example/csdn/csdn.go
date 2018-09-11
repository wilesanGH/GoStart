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
	"github.com/gocolly/colly/extensions"
	"time"
	"github.com/gocolly/colly/proxy"
	"github.com/satori/go.uuid"
	"GoStart/colly/example/csdn/orm"
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




	/********************************************************************************************************/
	//parentColly
	parentColly := colly.NewCollector(
		colly.AllowedDomains("www.csdn.net",
			"blog.csdn.net"),
		colly.Async(true),
		colly.DisallowedURLFilters(
			//regexp.MustCompile("https://blog.csdn.net/column.+"),
			regexp.MustCompile("https://blog.csdn.net/rss.h.+"),
			//regexp.MustCompile("https://blog.csdn.net/code/.+"),

		),

	)

	if p, err := proxy.RoundRobinProxySwitcher(
			"http://196.202.228.139:8080",
					"http://80.191.244.217:80",
					"http://94.74.177.181:80",
	); err == nil {
		parentColly.SetProxyFunc(p)
	}

	extensions.RandomUserAgent(parentColly)
	extensions.Referrer(parentColly)
	parentColly.Limit(&colly.LimitRule{
		RandomDelay:1*time.Second,
		DomainGlob:"*",
		Parallelism:2})


	subColly :=parentColly.Clone()

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
			subColly.Visit(reqUrl)
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



	//subColly.Limit(&colly.LimitRule{DomainGlob:"*",Parallelism:100})


	subColly.OnRequest(func(r *colly.Request) {

		//fmt.Println("subColly Visiting", r.URL.String())
	})
	subColly.OnHTML("div.blog-content-box", func(e *colly.HTMLElement) {

		csdnBlog := model.CSDN_BLOG{}
		csdnBlog.CsdnBase.Url = e.Request.URL.String()
		csdnBlog.Title = e.ChildText("h1.title-article")
		csdnBlog.Time = e.ChildText("span.time")
		csdnBlog.CsdnBase.Body = strings.Replace(strings.Replace(e.Text, "\n", "", -1), "\t", "", -1)
		csdnBlog.ReadCount = strings.Split(e.ChildText("span.read-count"), "：")[1]
		csdnBlog.Keywords = e.ChildAttrs("span.tags-box a", "data-track-view")
		//var keywordstring string
		e.ForEach("span.tags-box", func(_ int, el *colly.HTMLElement) {
			csdnBlog.Keywords = util.DeleteSpaceNTabForSlice(strings.Split(el.ChildText("a"), "\t"))
			//keywordstring = util.DeleteSpaceNTab(el.ChildText("a"))
		})

		csdnBlog.CsdnBase.Number = csdnBlog.CsdnBase.Url[strings.LastIndex(csdnBlog.CsdnBase.Url, "/")+1:]
		writer.Write([]string{
			csdnBlog.CsdnBase.Number,
			csdnBlog.Title,
			csdnBlog.CsdnBase.Url,
			fmt.Sprint(csdnBlog.Keywords)[:len(fmt.Sprint(csdnBlog.Keywords))-1][1:],
			csdnBlog.ReadCount,
			csdnBlog.CommentCount,
			csdnBlog.Time,
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

}


func GetCSDNBlog2() {


	count := 0
/*
	fName := "d:csdn_v3.csv"
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

*/




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
		colly.MaxDepth(3),

	)

	/*p, err := proxy.RoundRobinProxySwitcher(
		"http://196.202.228.139:8080",
		"http://80.191.244.217:80",
		"http://94.74.177.181:80",
		"https://54.39.97.250:3128",
		"socks5://127.0.0.1:1338",
	)
	if err != nil {
		fmt.Println(err)
	}
	parentColly.SetProxyFunc(p)
*/

	extensions.RandomUserAgent(parentColly)
	extensions.Referrer(parentColly)
	parentColly.Limit(&colly.LimitRule{
		RandomDelay:2*time.Second,
		})


	subColly :=parentColly.Clone()



	parentColly.OnRequest(func(r *colly.Request) {
		//fmt.Println("Visiting", r.URL.String())
	})

	parentColly.OnResponse(func(r *colly.Response) {
		reqUrl := r.Request.URL.String()
		if UrlRe.FindAllStringIndex(reqUrl, 1) == nil {
			return
		}
			subColly.Visit(reqUrl)
	})



	parentColly.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
			parentColly.Visit(e.Request.AbsoluteURL(link))
	})

	parentColly.OnError(func(r *colly.Response, e error) {
		fmt.Println("parent:Request URL：", r.Request.URL, "failed with response:", r, "\nError:", e)
	})


	subColly.OnRequest(func(r *colly.Request) {
		//fmt.Println("subColly Visiting", r.URL.String())
	})
	subColly.OnHTML("div.blog-content-box", func(e *colly.HTMLElement) {

		csdn_blog := model.CSDN_DETAIL{}
		u,_ := uuid.NewV4()
		csdn_blog.Id = u.String()
		csdn_blog.Url = e.Request.URL.String()
		csdn_blog.Title = e.ChildText("h1.title-article")
		csdn_blog.Time = e.ChildText("span.time")
		csdn_blog.Body = strings.Replace(strings.Replace(e.Text, "\n", "", -1), "\t", "", -1)
		csdn_blog.ReadCount = strings.Split(e.ChildText("span.read-count"), "：")[1]
		e.ForEach("span.tags-box", func(_ int, el *colly.HTMLElement) {
			csdn_blog.Keywords = strings.Replace(util.DeleteMoreTab(el.ChildText("a")),"\t"," ",-1)

		})

		csdn_blog.Number = csdn_blog.Url[strings.LastIndex(csdn_blog.Url, "/")+1:]

		//fmt.Printf("Insert csdbBlog:%v", csdn_blog)

		c,_ := orm.CsdnEngine.Table("csdnblog").Insert(&csdn_blog)
		fmt.Println(c)
		count++
		log.Printf("colly count: %d",count)



		/*log.Printf("count:%d\n", count)
		log.Println(csdn_blog)*/

	})

	subColly.OnError(func(r *colly.Response, e error) {
		fmt.Println("subColly:Request URL：", r.Request.URL, "failed with response:", r, "\nError:", e)
	})
	parentColly.Visit("http://www.csdn.net")




}