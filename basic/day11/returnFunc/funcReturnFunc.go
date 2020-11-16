package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 6, 4, 9}

	test1(nums)
	/*total := sum(nums...)
	fmt.Println(total)

	even := even(sum, nums...)
	fmt.Println(even)*/
}

func sum(nums ...int) int {
	var sum = 0
	for _, num := range nums {

		sum += num
	}

	return sum
}

func even(sum func(...int) int, nums ...int) int {
	return 0
}

func test1(nums ...[]int) {
	fmt.Println()
}
