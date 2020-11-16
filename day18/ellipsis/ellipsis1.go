package main

import "fmt"

func main() {
	var a = [...]int{1, 2, 3, 7, 8, 9} // 省略陣列長度讓編譯器自動計算。 同 a = [6]int{1, 2, 3, 7, 8, 9}
	fmt.Println(len(a))
	fmt.Println(a)
}

/* result:
6
[1 2 3 7 8 9]
*/
