package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("程式開始")

	p()
	time.Sleep(time.Second * 1)
	fmt.Println("主程式順利結束")
}

func p() {
	fmt.Println("即將發生空難...")
	panic("空難")
}

/* result:
程式開始
即將發生空難...
panic: 空難

goroutine 18 [running]:
main.p()
	/tmp/sandbox810050981/prog.go:18 +0x95
created by main.main
	/tmp/sandbox810050981/prog.go:11 +0x91
*/
