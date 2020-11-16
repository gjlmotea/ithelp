package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector()

	// 抓標籤Tag
	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	// 抓屬性值 AttrVal
	c.OnHTML("meta[name='description']", func(e *colly.HTMLElement) {
		fmt.Println(e.Attr("content")) // 抓此Tag中的name屬性 來找出此Tag，再印此Tag中的content屬性
	})

	// 抓類別Class 名稱
	c.OnHTML(".qa-list__title--ironman", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	// 抓唯一識別 ID
	c.OnHTML("#read_more", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.OnRequest(func(r *colly.Request) { // iT邦幫忙需要寫這一段 User-Agent才給爬
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	c.Visit("https://ithelp.ithome.com.tw/users/20125192/ironman/3155")
}
