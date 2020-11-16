package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func main() {
	var url = "https://member.ithome.com.tw/login"
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("1")
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("2")
	})

	c.OnResponseHeaders(func(r *colly.Response) {
		fmt.Println("3")
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("4")
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		fmt.Println("5")
	})

	c.OnXML("//footer", func(e *colly.XMLElement) {
		fmt.Println("6")
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("7")
	})

	c.Visit(url)
}
