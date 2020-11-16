package main

import "fmt"

type Animal interface {
	eat()
}

// 一開始看時，覺得寫interface之後再來引用宣告，
// 根本多此一舉，拿石頭砸自己的腳

type Cat struct {
	catName string
}

func (c Cat) eat() {
	fmt.Println("貓貓", c.catName, "開動哩")
}

func main() {
	var cat1 Animal = Cat{"肥貓一號"}
	cat1.eat()

	var cat2 Animal = Cat{"笨貓二號"}
	cat2.eat()

	// 哪天突然想不開，覺得貓太肥太醜，很不可愛
	// 想要棄貓養狗的時候
	// 只需實作"Dog"物件 及 "Dog.eat()" 的 func ↓↓↓
	var dog1 Animal = Dog{"開心狗一號"}
	dog1.eat()
}

type Dog struct {
	dogName string
}

func (d Dog) eat() {
	fmt.Println("狗狗", d.dogName, "不吃飯，只吃主人")
}
