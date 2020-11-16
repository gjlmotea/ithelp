package main

import "fmt"

func main()  {
	var r = []rune("0")
	fmt.Println(r)

	r = []rune("爆肝工程師的異世界安魂曲")
	fmt.Println(r)
}

/* result:
[48]
[29190 32925 24037 31243 24107 30340 30064 19990 30028 23433 39746 26354]
*/