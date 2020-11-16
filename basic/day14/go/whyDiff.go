package main

import (
	"fmt"
	"runtime"
	"time"
)

var c1, c2 int64

func main() {
	runtime.GOMAXPROCS(1)
	go p1()
	go p2()
	time.Sleep(time.Millisecond * 1000)

	fmt.Println(c1, c2)
}

func p1() {
	for {
		c1++
	}
}

func p2() {
	for i := 0; i < 100000000000; i++ {
		c2++
	}
}
