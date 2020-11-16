package main

import (
	"fmt"
	"time"
)

var count = 0

func main() {
	// runtime.GOMAXPROCS(1)  // 只讓一個線程運作就能解決問題。但...想要更快，就是要多核心嘛！單核心怎麼星爆？
	for i := 0; i < 10000; i++ {
		go race()
	}
	time.Sleep(time.Millisecond * 100)
	fmt.Println(count)
}

func race() {
	count++
}

/* result:
9763
*/
