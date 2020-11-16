package main

import "fmt"

const PI = 3.14159
const HELLO = "Hello World"

const (
	A = iota       // 0
	B              // 1
	_              // 2 佔位符也會被計算
	C              // 3
	D = iota * 0.1 // 0.4 接續前面的 iota
	E              // 5
	F              // 6
	G              // 7
)

const (
	X = iota + 100 // 100
	Y              // 101
	Z              // 102
)

const (
	b1 = 1 << iota // 1  右側被塞入0個bit (2^0 二的零次方)
	b2             // 2  右側被塞入1個bit (2^1 二的一次方)
	b3             // 4  右側被塞入2個bit
	b4             // 8
	b5             // 16
)

func main() {
	fmt.Println(HELLO, PI)
	fmt.Println(A, B, C, D, E, F, G)
	fmt.Println(X, Y, Z)
	fmt.Println(b1, b2, b3, b4, b5)
}

/* result:
Hello World 3.14159
0 1 3 0.4 0.5 0.6 0.7
100 101 102
1 2 4 8 16
*/
