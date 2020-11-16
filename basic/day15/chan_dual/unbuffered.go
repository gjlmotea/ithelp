package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func1(ch)
	ch <- 100

	go func2(ch)
	got := <-ch
	fmt.Println(got)
}

func func1(ch chan int) {
	i := <-ch
	fmt.Println(i)
}

func func2(ch chan int) {
	time.Sleep(time.Second * 2)
	ch <- 999
}

/* result:
100
999
*/
