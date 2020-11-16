package main

import (
	"log"
	"time"
)

func main() {
	log.SetFlags(5)
	ch := make(chan int, 5)
	go func4(ch)
	for i := 0; i < 10; i++ {
		ch <- i
		log.Println("main sent", i)
	}
	time.Sleep(time.Second)
}

func func4(ch chan int) {
	for {
		i := <-ch
		log.Println("func got", i)
		time.Sleep(time.Millisecond * 100)
	}
}

/* result:
func got 0
main sent 0
main sent 1
main sent 2
func got 1
main sent 3
func got 2
main sent 4
func got 3
main sent 5
func got 4
main sent 6
func got 5
main sent 7
func got 6
main sent 8
func got 7
main sent 9
func got 8
func got 9
*/
