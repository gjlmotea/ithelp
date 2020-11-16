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
	*Bag   // 你中有針
}

type Bag struct {
	Wallet
	PencilBox
	Books string
}

func main() {
	var bag = Bag{
		Wallet{Card: "世華泰國信用無底洞卡", Red100: 5},
		PencilBox{Pen: "Cross", Pencil: "Pentel"},
		"Go繁不及備載",
	}
	bag.PencilBox.Bag = &bag // 包包裡放針

	fmt.Printf("%+v", *bag.PencilBox.Bag)
}

/* result:
{Wallet:{Blue1000:0 Red100:5 Card:世華泰國信用無底洞卡} PencilBox:{Pencil:Pentel Pen:Cross Bag:0xc000046060} Books:Go繁不及備載}
*/
