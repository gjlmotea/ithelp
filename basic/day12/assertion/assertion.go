package main

import "fmt"

func main() {
	var hello interface{} = "hello"
	helloStr := hello.(string)
	fmt.Println(helloStr)

	helloInt, ok := hello.(int)
	fmt.Println(helloInt, ok)
}

/* result:
hello
0 false
*/
