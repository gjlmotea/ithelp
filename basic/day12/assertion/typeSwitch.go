package main

import (
	"errors"
	"fmt"
)

func main() {
	whatType(nil)
	whatType(123)
	whatType("Hello")
	whatType(true)
	whatType(errors.New("something went wrong"))
}

func whatType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case nil:
		fmt.Println("nil")
	case error:
		fmt.Println("error")
	default:
		fmt.Println("unknown", v)
	}
}
