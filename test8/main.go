package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	// 创建一个新的爬虫实例
	c := colly.NewCollector()

	// 在访问每个页面成功时调用此回调函数
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response received", r.StatusCode)
		// 在页面内容中查找演唱会门票信息
		r.ForEach("li.item", func(_ int, e *colly.HTMLElement) {
			fmt.Printf("%s %s\n", e.ChildText(".left em"), e.ChildText(".title a"))
		})
	})

	// 访问大麦网的演唱会门票页面
	c.Visit("https://search.damai.cn/search.htm?spm=a2oeg.home.top.ditem_1_1b66f537.KG9P0d&ctl=%E6%BC%94%E5%94%B1%E4%BC%9A&order=1&cty=&sctl=&tsg=0")
}
