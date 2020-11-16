package main

import "fmt"

func main() {
	slice := []int{2, 3, 5}

	sum1 := sumUnpacking(slice...) // 把slice 解開、剝皮後傳入，同下
	fmt.Println(sum1)

	sum2 := sumUnpacking(2, 3, 5) // 可變參數函式
	fmt.Println(sum2)

	sum3 := sumSlice(slice) // 不曉得int長度，也可以直接包成一個slice型別來傳遞
	fmt.Println(sum3)
}

func sumUnpacking(nums ...int) int { // 傳入int但不曉得參數長度為何
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}

func sumSlice(nums []int) int { // 傳入slice
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}

/* result:
10
10
10
*/
