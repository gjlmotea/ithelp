package main

import "fmt"

var x = 123 // 全域變數

func main() {
	fmt.Println(&x)

	var x = 123 // 在`func區塊`。全新。
	fmt.Println(&x)

	// var x, y = 123, 100 // 因為x已經宣告過了，不能再用`var`。但以下卻可以
	x, y := 123, 100 // 在`func區塊`。此時x是上面宣告過的區域變數，而y是新的
	fmt.Println(&x, &y)

	if true {
		var x, y = 123, 100 // 在新區域`if區塊`。此時x跟y都是全新的。
		// 可以換成短變數宣告。

		fmt.Println(&x, &y)
	}

	fmt.Println(&x) // 脫離if區域，回到`func區塊`
}

/* 運行結果
0x11662a0
0xc000016080
0xc000016080 0xc000016088
0xc000016090 0xc000016098
0xc000016080
*/
