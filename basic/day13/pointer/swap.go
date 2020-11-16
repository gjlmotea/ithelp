package main

import "fmt"

func main() {
	a, b := 10, 20
	swap(a, b)
	fmt.Println(a, b)

	Swap(&a, &b)
	fmt.Println(a, b)
}

// 我以為有用的swap function
func swap(a int, b int) {
	temp := a
	a = b
	b = temp
}

// 實際上真正有用的Swap function
func Swap(a *int, b *int) {
	//fmt.Println(a, b) //0xc00000a108 0xc00000a120
	temp := *a
	*a = *b
	*b = temp
}
