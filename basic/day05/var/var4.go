package main

import "fmt"

var i = 123 // 全域變數

func main() {
	fmt.Println(&i)

	i = 123
	fmt.Println(&i)

	i := 123 // 全新的。func區塊內的區域變數
	fmt.Println(&i)

	i, j := 123, 100
	fmt.Println(&i, j) // 奇怪？i沒有變新的啊
}

/* 運行結果
0x11662a0
0x11662a0
0xc0000b2008
0xc0000b2008 100
*/
