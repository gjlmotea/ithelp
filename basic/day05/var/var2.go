package main

import (
	"fmt"
)

func main()  {
	var b int = 123
	a, b := 100, 99
	c, b := 0, 1
	fmt.Println(a, b, c)
}

// result:
// 100 1 0
