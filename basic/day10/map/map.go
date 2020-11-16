package main

import (
	"fmt"
)

func main() {
	var Male = map[bool]string{
		true:  "公",
		false: "母",
	}

	var Number map[string]int // 沒有{}賦值、也沒有make實體空間，所以為 nil
	fmt.Println(Number == nil)
	Number = make(map[string]int) // 用make創造實體空間
	Number["零"] = 0
	Number["壹"] = 1
	Number["貳"] = 2
	Number["參"] = 3

	//--------

	var Size = map[string]string{
		"big":    "大",
		"medium": "中",
		"small":  "小",
	}
	fmt.Print(Male[true], Number["參"])

	//--------

	fmt.Println()
	for key, value := range Size {
		fmt.Println(key, value)
	}
}

/* result:
true
公3
big 大
medium 中
small 小
*/
