package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func main() {
	// spiderV1()

	c := colly.NewCollector()

	// 在每个请求之前，打印请求的URL
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// 在响应中查找列表信息
	c.OnHTML(".el-card__body", func(e *colly.HTMLElement) {
		// 使用GoQuery解析HTML，获取所需信息
		doc := e.Response.Body
		dom, err := goquery.NewDocumentFromReader(bytes.NewBuffer(doc))
		if err != nil {
			fmt.Println(err)
			return
		}
		title := dom.Find(".name").Text()
		avatar := dom.Find(".avatar img").AttrOr("src", "")

		fmt.Printf("Title: %s\n", title)
		fmt.Printf("Avatar: %s\n", avatar)
	})

	// 发起请求
	c.Visit("https://ssr1.scrape.center/")

	fmt.Println("Finished")
}
