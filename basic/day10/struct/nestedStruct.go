package main

import "fmt"

type Wallet struct {
	Blue1000 int // 藍色小朋友
	Red100   int // 紅色國父
	Card     string
}

type PencilBox struct {
	Pencil string
	Pen    string
}

type Bag struct {
	Wallet    // 直接放入結構就好
	PencilBox // 直接放入結構就好
	Books     string
}

type Person struct {
	Bag  // 放Bag這個物件
	Name string
}

func main() {
	var bag = &Bag{
		Wallet{Card: "世華泰國信用無底洞卡", Red100: 5},
		PencilBox{Pen: "Cross", Pencil: "Pentel"},
		"Go繁不及備載",
	}

	var Tommy = Person{}
	Tommy.Name = "Tommy"
	Tommy.Bag = *bag // 取出bag位址裡面的東西

	fmt.Printf("%+v", Tommy)
}

/* result:
{Bag:{Wallet:{Blue1000:0 Red100:5 Card:世華泰國信用無底洞卡} PencilBox:{Pencil:Pentel Pen:Cross} Books:Go繁不及備載} Name:Tommy}
*/
