package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// link := e.Attr("href")
		fmt.Printf("e %+v\n", e)

		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		// c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://ssr1.scrape.center/")
}
