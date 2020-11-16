package main

import "fmt"

func main() {
	str := d()
	fmt.Println(str)

	x, y := f()
	fmt.Println(x, y)

	g()
}

func d() (e string) {
	e = "I am e!"
	return
}

func f() (int, int) {
	return 1 + 1, 2 + 2
}

func g() (s int, t int) {
	t = 100 * 100
	fmt.Println("ggg")
	return 10 * 10, t
}

/* result:
I am e!
2 4
ggg
*/
