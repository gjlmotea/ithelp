package main

import "fmt"

func main() {
	assign1(50)

	assign2(50)
}

func assign1(a int) int {
	a = 100              // 老闆更動了a為100
	defer fmt.Println(a) // 任務交代下來的時候 a值是100，然後盡可能地拖延
	return a
}

func assign2(a int) int {
	defer fmt.Println(a) // 任務交代下來的時候 a值是50，然後盡可能地拖延
	a = 100              // 老闆更動了a為100
	return a
}

/* result:
100
50
*/
