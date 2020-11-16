package main

import "fmt"

func main() {
	var a int = 10
	var b *int
	b = &a
	fmt.Println(a, b)

	var c string = "hi"
	var d *string
	d = &c
	*d = "who" // 透過`*向址取值`的方式來改變變數裡面的內容值。
	fmt.Println(c, d)
}

/* result:
10 0xc00000a108
who 0xc0000401f0
*/
