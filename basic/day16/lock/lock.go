package main

import (
	"fmt"
	"sync"
	"time"
)

var count = 0
var m sync.Mutex

func main() {
	for i := 0; i < 10000; i++ {
		go race()
	}
	time.Sleep(time.Millisecond * 100)
	fmt.Println(count)
}

func race() {
	m.Lock()
	count++
	m.Unlock()
}

/* result:
10000
*/
