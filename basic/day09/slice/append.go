package main

import "fmt"

func main() {
	x := []int{1, 2, 3} //看到輸出正常，豪開心呀
	x = append(x, 4, 5, 6)
	fmt.Println(x)

	//--------

	a := make([]int, 0, 10) //給足夠容量
	b := append(a, 1, 2, 3)
	_ = append(a, 99, 88, 77)
	fmt.Println(b)

	//--------

	aa := make([]int, 0, 2) //給超小容量
	bb := append(aa, 1, 2, 3)
	_ = append(aa, 99, 88, 77)
	fmt.Println(bb)
}

/* result:
[1 2 3 4 5 6]
[99 88 77]
[1 2 3]
*/
