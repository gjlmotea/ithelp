package main

import "fmt"

func main() {
	nums := []int{100, 99, 98}
	for index, num := range nums {
		fmt.Println(index, num)
	}
	for index := range nums {
		fmt.Println(index)
	}

	//--------

	fruits := map[string]string{"a": "apple", "b": "banana"}
	for index, fruit := range fruits {
		fmt.Println(index, fruit)
	}
	for index := range fruits {
		fmt.Println(index)
	}
}

/* result:
0 100
1 99
2 98
0
1
2
a apple
b banana
a
b
*/
