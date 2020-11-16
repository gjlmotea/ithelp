package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

var token string

func main() {
	var url = "https://member.ithome.com.tw/login"
	c := colly.NewCollector()

	// 拿到這次登入的Token
	c.OnHTML("input[name='_token']", func(e *colly.HTMLElement) {
		token = e.Attr("value")
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.Headers.Get("Set-Cookie"))
	})
	c.Visit(url)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Chrome/84.0.4147.89 Safari/537.36")
		r.Headers.Set("Host", "https://member.ithome.com.tw")
		r.Headers.Set("Origin", "https://member.ithome.com.tw")
		r.Headers.Set("Referer", "https://member.ithome.com.tw/login")
		// 這幾行在這iT邦幫忙沒有起到作用，但有些網站會照這些資訊判斷、阻擋其他來源
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(string(r.Body))
	})

	var formData = map[string]string{
		"account":  "collytest",
		"password": "collytest",
		"_token":   token,
	}
	err := c.Post(url, formData) // 進到該url 執行POST
	if err != nil {
		log.Println(err)
	}

}
