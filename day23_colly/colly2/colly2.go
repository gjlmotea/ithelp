package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

var count = 0

func main() {

	c := colly.NewCollector()

	// 當Visit訪問網頁後，在網頁響應(Response)之後、發現這是HTML格式 執行的事情
	c.OnHTML(".qa-list__title-link", func(e *colly.HTMLElement) { // 每找到一個符合 goquerySelector字樣的結果，便會進這個OnHTML一次
		fmt.Println(e.Text)
		count++
	})

	c.OnRequest(func(r *colly.Request) { // iT邦幫忙需要寫這一段 User-Agent才給爬
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	c.Visit("https://ithelp.ithome.com.tw/users/20125192/ironman/3155")

	fmt.Println(count) // count值為10，代表原始碼中 有10個符合規則的結果，總共進了OnHTML func 10次
}
