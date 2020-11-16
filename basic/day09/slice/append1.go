package main

import "fmt"

func main() {
	a := make([]int, 0, 10)
	b := append(a, 1, 2, 3)
	_ = append(a, 99, 88, 77)
	fmt.Println(b)

	aa := make([]int, 0, 2)
	bb := append(aa, 1, 2, 3)
	_ = append(aa, 99, 88, 77)
	fmt.Println(bb)
}
