package main

import (
	"fmt"
	"time"
)

func main() {
	go test()
	time.Sleep(time.Second * 1)
}

func test() {
	fmt.Println("嚇嚇叫")
}

/* result:
嚇嚇叫
*/
