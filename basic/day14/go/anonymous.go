package main

import (
	"fmt"
	"time"
)

func main() {
	go func() { // 把這裡的go併發去掉，就變成匿名函式了
		for i := 0; i < 10000; i++ {
			fmt.Print("匿名併發函式 ")
		}
	}() // 這邊有小括號有點奇怪對不？ 因為他是一個函式，在呼叫時需要()

	for i := 0; i < 10000; i++ {
		fmt.Print("main ")
	}

	time.Sleep(10000)
}

/* result:
main main main main main main main main main main main 匿名併發函式 匿名併發函式 匿名併發函式 匿名併發函式...
*/
