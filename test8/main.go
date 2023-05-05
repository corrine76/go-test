package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"log"
	"os"
	"test8/spider"
)

func main() {
	spider.V3()
}

func run() {
	// 文件暂存
	fName := "go-spider-test-v1.json"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()

	movies := make([]spider.Movie, 0, 200)
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

		movie := spider.Movie{
			Name:  title,
			Cover: avatar,
		}
		movies = append(movies, movie)
	})

	// 发起请求
	c.Visit("https://ssr1.scrape.center/")

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	enc.Encode(movies)
	fmt.Println("Finished")
}
