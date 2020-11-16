package main

import "fmt"

func main() {
	s1 := sum()
	fmt.Println(s1)

	s2 := sum(1, 5, 9)
	fmt.Println(s2)
}

func sum(nums ...int) int { // 函式中的可變長度參數。這裡的 nums可視為 slice[]
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}

/* result:
0
15
*/
