package main

import "fmt"

type Programmer struct {
	Question string
}

func main() {
	var I Programmer // 物件實體化方式之一
	I.Question = "Hey, Can I Sleep?"
	I.Ask()

	var boss = Boss{} // 物件實體化方式之一
	boss.Reply()
}

func (CIS *Programmer) Ask() {
	question := CIS.Question
	fmt.Println(question)
}

type Boss struct {
}

func (B *Boss) Reply() {
	fmt.Println("No 啦！")
}

/* result:
Hey, Can I Sleep?
No 啦！
*/
