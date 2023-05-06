package spider

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Concert struct {
	Title        string
	Desc         string
	Poster       string
	Price        string
	Like         string
	View         string
	Addr         string
	Date         string
	Artists      []string
	DiscountCode string
}

func V3() {
	// 使用json文件存储数据
	fname := "go-spider-test-v3.json"
	file, err := os.Create(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	concerts := make([]Concert, 0)
	c := colly.NewCollector()

	// todo 获取到热门城市id
	// <span class="hot-item single-city" data-citypinyin="beijing" data-cityname="北京" data-cityid="1101">北京</span>
	// 遍历热门城市演唱会列表

	page := 1
	for {
		c.OnRequest(func(r *colly.Request) {
			fmt.Println("visiting", r.URL.String())
		})

		c.OnHTML("a.show-items", func(e *colly.HTMLElement) {
			doc := e.DOM
			title := doc.Find("div.show-name").AttrOr("title", "")
			desc := doc.Find("div.show-desc").Text()
			poster := doc.Find("div.show-poster img").AttrOr("data-src", "")
			addr := doc.Find("div.show-addr").Text()
			date := doc.Find("div.show-time").Text()
			like := strings.Replace(doc.Find("i.icon-like").Parent().Text(), " ", "", -1)
			view := strings.Replace(doc.Find("i.icon-scan").Parent().Text(), " ", "", -1)
			code := RandStringBytesMaskImprSrc(8) // 生成8位随机码，包含英文和数字

			if addr == "what" {
				addr = "待定"
			}
			like = strings.TrimSpace(like)
			view = strings.TrimSpace(view)
			fmt.Println(title, desc, poster, addr, date, like, view, code)

			concerts = append(concerts, Concert{
				Title:        title,
				Desc:         desc,
				Poster:       poster,
				Price:        "1元起",
				Like:         like,
				View:         view,
				Addr:         addr,
				Date:         date,
				Artists:      []string{},
				DiscountCode: code,
			})
		})

		c.Visit(fmt.Sprintf("https://www.moretickets.com/list/3101-concerts/p%d", page))
		page++
		if page > 4 {
			break
		}
	}

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	enc.Encode(concerts)
}

// RandStringBytesMaskImprSrc 生成8位随机码，包含英文和数字
func RandStringBytesMaskImprSrc(length int) string {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 定义随机码包含的字符
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	// 生成随机码
	code := make([]rune, length)
	for i := range code {
		code[i] = letters[rand.Intn(len(letters))]
	}
	return string(code)
}
