package main

func main() {
	ch := make(chan int, 2)
	go func3(ch)
	ch <- 100
	ch <- 99

	ch <- 98 // 發生deadlock
}

func func3(ch chan int) {
}

/* result:
fatal error: all goroutines are asleep - deadlock!
*/
