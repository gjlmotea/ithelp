package main

import (
	"fmt"
	"log"
)

func main() {
	for i := 1; i <= 3; i++ {
		log.Println(i)
		fmt.Println(i)
		//time.Sleep(time.Millisecond * 10)
		log.Println("hi")
		fmt.Println("hi")
	}

}

/* result: (in Win10)
2020/09/13 14:42:20 1
2020/09/13 14:42:20 hi
2020/09/13 14:42:20 2
2020/09/13 14:42:20 hi
2020/09/13 14:42:20 3
2020/09/13 14:42:20 hi
1
hi
2
hi
3
hi
*/
