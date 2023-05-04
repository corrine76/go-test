package spider

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"log"
	"os"
)

func V2() {
	// 文件暂存
	fName := "go-spider-test-v2.json"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()

	movies := make([]Movie, 0)
	c := colly.NewCollector()

	c.OnHTML(".el-card__body", func(e *colly.HTMLElement) {
		// 获取电影名称
		name := e.ChildText(".name h2")

		// 获取所有分类标签
		var categories []string
		e.DOM.Find(".categories .category span").Each(func(_ int, s *goquery.Selection) {
			categories = append(categories, s.Text())
		})

		// 获取电影时长
		duration := e.ChildText(".info:nth-child(2)")

		// 获取上映日期
		releaseDate := e.ChildText(".info:nth-child(3)")

		// 获取评分
		score := e.ChildText(".score")

		// 获取详情页URL
		href, _ := e.DOM.Find(".name a").Attr("href")
		detailURL := "https://example.com" + href // example.com替换为实际网站的域名

		movie := Movie{}
		movie.Name = name
		movie.Categories = categories
		movie.URL = detailURL
		movie.Duration = duration
		movie.ReleaseDate = releaseDate
		movie.Score = score

		// todo 实现cover和desc的爬取

		movies = append(movies, movie)
	})

	c.Visit("https://ssr1.scrape.center/page/1")

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	enc.Encode(movies)
}
