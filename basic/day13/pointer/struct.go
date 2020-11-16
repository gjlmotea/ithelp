package main

import "fmt"

type Cat struct {
	catName string
}

func (c Cat) eat() {
	fmt.Println("貓貓", c.catName, "開動哩")
}

func (c Cat) rename(newName string) {
	c.catName = newName
}

func main() {
	var cat1 = Cat{"肥貓一號"}
	cat1.eat()

	var cat2 = Cat{"笨貓二號"}
	cat2.rename("聰明貓三號") // 奇怪，怎麼改名失敗了
	cat2.eat()
}
