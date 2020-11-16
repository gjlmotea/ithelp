package main

import (
	"fmt"
	"time"
)

var counter int

func main() {
	//runtime.GOMAXPROCS(1)
	count1()

	time.Sleep(time.Second * 1)
	fmt.Println(counter)
}

func count1() {
	counter++
	go count1()
}

func count2() {
	counter++
	go count2()
}
