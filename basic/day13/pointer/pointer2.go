package main

import "fmt"

func main() {
	x := 1
	p := &x          //p(type: *int)指向x
	fmt.Println(x)   //1
	fmt.Println(p)   //p指向的位址
	fmt.Println(*p)  //p指向的位址的值，意即變數x
	fmt.Println(&p)  //p本身的位址
	y := &p          //y存放了 p本身的位址
	fmt.Println(**y) //到y取值(到p本身的位址取值) 再取值，意即變數x
	**y = 100
	fmt.Println(x)
}

/* result:
1
0xc00011e090
1
0xc000148018
1
100
*/
