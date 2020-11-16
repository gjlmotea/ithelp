package main

import (
	"fmt"
)

func main() {
	var a [5]int
	a[0] = 10
	a[1] = 100
	a[2] = 1000
	fmt.Println(a)

	b := [5]int{1, 2, 3}
	fmt.Println(b)

	c := [5]int{
		10,
		20,
		30,
		55, //使用多行宣告的話，最後一個元素要逗號
	}
	fmt.Println(c[1:3])

	d := [...]int{4, 6, 8} //用...省略符號，讓go判斷長度
	fmt.Println(d)
}

/* result:
[10 100 1000 0 0]
[1 2 3 0 0]
[20 30]
[4 6 8]
*/
