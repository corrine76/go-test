package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now().Format("2006-01-02")
	fmt.Println("当前时间", now)

	// 发起 HTTP GET 请求获取网页内容
	resp, err := http.Get("http://top.baidu.com/buzz?b=1&fr=topindex&t=" + now + "1200")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 使用 goquery 解析网页内容
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("接收响应", doc)

	// 使用正则表达式获取热搜词
	r := regexp.MustCompile(`^(\d+)\s+(\S+)\s+(\S+)\s+(\S+)\s*$`)
	doc.Find(".keyword .list-title").Each(func(i int, s *goquery.Selection) {
		fmt.Sprintln("开始匹配", s.Text())
		text := strings.TrimSpace(s.Text())
		matches := r.FindStringSubmatch(text)
		if len(matches) == 5 {
			rank, _ := strconv.Atoi(matches[1])
			fmt.Printf("%d. %s\n", rank, matches[2])
		}
	})
}
