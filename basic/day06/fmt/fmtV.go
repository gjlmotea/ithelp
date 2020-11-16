package main

import "fmt"

type Name struct {
	A string
	B bool
	C int
}
func main()  {
	fmt.Printf("%v	\n", Name{})
	fmt.Printf("%+v	\n", Name{})
	fmt.Printf("%#v	\n", Name{})
}

// result:
// { false 0}
// {A: B:false C:0}
// main.Name{A:"", B:false, C:0}