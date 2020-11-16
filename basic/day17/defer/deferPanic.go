package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		fmt.Println("我很懶，想等到退出func的時候再印東西")
	}()
	os.Exit(0) // func直接被砍了
}

/* result:

 */
