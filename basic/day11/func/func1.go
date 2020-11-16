package main

import "fmt"

func main() {
	a()
	b := b()
	fmt.Println(b)
	fmt.Println(c() + 6)
}

func a() {
	fmt.Println("Hello")
}

func b() string {
	return "Hi"
}

func c() int {
	return 1 + 5
}

/*
Hello
Hi
12
*/
