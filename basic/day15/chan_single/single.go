package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(<-chan int)
	test(ch)
}

func test(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println("func got", i)
		time.Sleep(time.Millisecond * 100)
	}
}

/* result:
 */
