package main

import "fmt"

var a = 50

func main() {
	defer fmt.Println("退出main才執行")
	fmt.Println("Hi")

	defer add_a()

	fmt.Println(a)
}

func add_a() {
	a += 100
}

/* result:
Hi
50
退出main才執行
*/
