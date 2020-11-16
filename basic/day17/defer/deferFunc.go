package main

import "fmt"

func main() {
	fmt.Println(func1())
	fmt.Println(func2())
}

func func1() int {
	var a int
	defer func() {
		a = 100
	}()
	return a // defer：『喔 要回傳a了喔，可是func還沒退出所以我不想做事，反正回上司也沒有規定要回傳哪個a，所以擺爛。』
}

func func2() (a int) {
	defer func() {
		a = 100
	}()
	return a //defer：『蛤，要回傳了喔？雖然想擺爛，但上司一開始指名規定要回傳a，先趕一下進度好了。』
}

/* result:
0
100
*/
