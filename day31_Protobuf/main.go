package main

import (
	"fmt"

	pb "github.com/gjlmotea/ithelp/day31_Protobuf/proto"
)

func main() {
	teacher := pb.Teacher{Name: "Jack", Age: 32}
	fmt.Printf("%+v\n", teacher)

	tName := teacher.GetName()
	tAge := teacher.GetAge()
	fmt.Println(tName, tAge)
}
