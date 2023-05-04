package spider

import (
	"encoding/json"
	"github.com/gocolly/colly"
	"log"
	"os"
)

func V1() {
	// 文件暂存
	movies := make([]Movie, 0, 200)
	fName := "go-spider-test.json"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()

	c := colly.NewCollector(
	// colly.AllowedDomains("https://ssr1.scrape.center/"),
	// colly.CacheDir("./coursera_cache"),
	)

	// On every <a> element which has "href" attribute call callback
	c.OnHTML("div.el-row", func(e *colly.HTMLElement) {
		if e.Attr("class") == "el-col el-col-24 el-col-xs-8 el-col-sm-6 el-col-md-4" {
			return
		}
		log.Printf("%+v\n", e)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	c.Visit("https://ssr1.scrape.center/")

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	enc.Encode(movies)
}
