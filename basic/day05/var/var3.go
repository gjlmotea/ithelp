package main

import "fmt"

var a = "Hello!"

func main()  {
	b := 10
	fmt.Println(&a, a, &b, b)

	a, b := 100, 99
	fmt.Println(&a, a, &b, b)
}

/* result:
0x54dc50 Hello! 0xc00002c008 10
0xc00002c048 100 0xc00002c008 99
*/

