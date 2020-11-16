package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

var wordCount = 0
var chineseCount = 0

func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) { // iT邦幫忙需要寫這一段 User-Agent才給爬
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	c.OnHTML(".qa-list__title-link", func(e *colly.HTMLElement) {
		// fmt.Println(e.Text, e.Attr("href"))
		// e.Text 印出 <a> tag 裡面的文字，也就是文章標題
		// e.Attr("href") 則是找到 <a> tag裡面的 href元素

		linksStr := e.Attr("href")
		linksStr = strings.Replace(linksStr, " ", "", -1) // 把空白以""取代掉
		links := strings.Split(linksStr, "\n")            // 以換行符號(\n)做為分隔來切割字串

		for _, link := range links {
			c2 := colly.NewCollector() // 這邊要在迴圈一開始再宣告一個 Collector，才不會與原本的混合到
			c2.OnHTML(".qa-markdown", func(e2 *colly.HTMLElement) {
				fmt.Println(e2.Text) // 印出 qa-markdown class中的文字（Go繁不及備載 文章的內文）

				countWord(e2.Text)

			})
			c2.OnRequest(func(r *colly.Request) { // iT邦幫忙需要寫這一段 User-Agent才給爬
				r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
			})
			c2.Visit(link) // 找到<a>連結網址後，點進去訪問
		}
	})

	c.Visit("https://ithelp.ithome.com.tw/users/20125192/ironman/3155?page=1")
	c.Visit("https://ithelp.ithome.com.tw/users/20125192/ironman/3155?page=2")
	c.Visit("https://ithelp.ithome.com.tw/users/20125192/ironman/3155?page=3")

	fmt.Println("英文+中文+數字 共", wordCount, "字")
	fmt.Println("純中文字 共", chineseCount, "字")
}

func countWord(input string) {
	for _, word := range input {
		if word != 32 && word != 10 { // 計算有多少非空白(space)以及換行(\n)的字數
			wordCount++
		}
		if word > 256 { // 計算有多少中文字數(編碼比ASCII大的字)
			chineseCount++
		}
	}
}
