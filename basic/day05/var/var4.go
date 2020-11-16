package main

import "fmt"

var a = 123

func main()  {
	fmt.Println(&a)

	a = 123
	fmt.Println(&a)

	a := 123
	fmt.Println(&a)

	a, b := 100, 99
	fmt.Println(&a)

	a, c := 456, "hello"
	fmt.Println(a, b, c)
}

/* result:
0x53f108
0x53f108
0xc000094018
0xc000094018
456 99 hello
*/

