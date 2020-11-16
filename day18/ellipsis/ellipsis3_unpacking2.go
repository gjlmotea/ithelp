package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8}

	// 將兩 Slice 合併(append) 的方法 - 1
	a1 := append(slice1, slice2[0], slice2[1], slice2[2]) // append用法，每個參數只能附上一個int
	fmt.Println(a1)

	// 將兩 Slice 合併(append) 的方法 - 2
	a2 := slice1
	for _, num := range slice2 { // 如方法1，只是這次透過for迴圈來迭代完成
		a2 = append(a2, num)
	}
	fmt.Println(a2)

	// 將兩 Slice 合併(append) 的方法 - 3
	a3 := append(slice1, slice2...) // 直接將slice2 剝皮解壓縮後(Unpacking)再執行append，取代上面使用For迭代的方法
	fmt.Println(a3)
}

/* result:
[1 2 3 4 5 6 7 8]
[1 2 3 4 5 6 7 8]
[1 2 3 4 5 6 7 8]
*/
