package main

import "fmt"

type cat struct {
	name string
}

func main() {
	var c = &cat{name: "始祖貓"}
	fmt.Println(c, &c)

	n1 := newCat("")
	n2 := newCat("複製貓三號")

	fmt.Println(n1, &n1)
	fmt.Println(n2, &n2)

	var c2 = new(cat) // 內建的new方法
	fmt.Println(c2, &c2)
}

func newCat(n string) *cat {
	return &cat{name: n}
}

/* result:
&{始祖貓} 0xc0000ca018
&{} 0xc0000ca028
&{複製貓三號} 0xc0000ca030
&{} 0xc0000ca038
*/
