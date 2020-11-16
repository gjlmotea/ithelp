package main

import "fmt"

func main() {
	var i int = 2
	switch i {
	case 1:
		fmt.Println("i is 1")
	case 2:
		fmt.Println("i is 2")
	case 3:
		fmt.Println("i is 3")
	default:
		fmt.Println("i is not 1, 2, 3")

	}
	//--------
	var j int = 10
	switch {
	case j == 1:
		fmt.Println("j is 1")
	case j == 2:
		fmt.Println("j is 2")
	case j == 3:
		fmt.Println("j is 3")
	default:
		fmt.Println("j is not 1, 2, 3")
	}
}

// result:
// a is 10
