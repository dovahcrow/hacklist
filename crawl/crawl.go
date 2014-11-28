package main

import (
	"fmt"
	"github.com/PuerkitoBio/gocrawl"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/http/cookiejar"
	"time"
)

var titles []string

func init() {
	jar, _ := cookiejar.New(nil)
	gocrawl.HttpClient.Jar = jar
}

type myextender struct {
	gocrawl.DefaultExtender
}

func (this *myextender) Visit(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (harvested interface{}, findLinks bool) {

	title, _ := doc.Find("title").Html()
	titles = append(titles, title)
	return nil, true
}
func main() {

	gc := gocrawl.NewCrawler(&myextender{})
	gc.Options.CrawlDelay = 10 * time.Nanosecond
	gc.Options.MaxVisits = 10
	gc.Options.LogFlags = gocrawl.LogInfo
	gc.Options.UserAgent = "aloha"
	err := gc.Run("http://www.baidu.com/")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(titles)
}
