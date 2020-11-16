package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("你各位啊，現在開始休息，三秒鐘後記得回來。")

	wg := sync.WaitGroup{} // 也可以var wg = sync.WaitGroup{}，或者不要實體化 var wg sync.WaitGroup
	wg.Add(2)              // 總共有兩位新兵

	go rest(&wg)
	go rest(&wg)

	fmt.Println("===你各位再慢慢來沒關係啊===")
	wg.Wait()
	fmt.Println("===集合完畢===")
}

func rest(wg *sync.WaitGroup) {
	time.Sleep(time.Second * 3)
	fmt.Println("新兵休息完畢。")
	wg.Done() // 跑去集合
}

/* result:
你各位啊，現在開始休息，三秒鐘後記得回來。
===你各位再慢慢來沒關係啊===
新兵休息完畢。
新兵休息完畢。
===集合完畢===
*/
