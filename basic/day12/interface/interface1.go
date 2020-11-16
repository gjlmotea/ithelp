package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{}
	fmt.Println(a, reflect.TypeOf(a))
	a = 123
	fmt.Println(a, reflect.TypeOf(a))
	a = "hi"
	fmt.Println(a, reflect.TypeOf(a))
	a = true
	fmt.Println(a, reflect.TypeOf(a))
}
