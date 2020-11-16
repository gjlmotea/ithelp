package main

import "fmt"

func main() {
	defer print1()
	defer print2()
}

func print1() {
	fmt.Println("p1")
}

func print2() {
	fmt.Println("p2")
}

/* result:
p2
p1
*/
