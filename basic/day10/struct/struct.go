package main

import "fmt"

type Res struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

func main() {
	res1 := new(Res)

	var res2 = new(Res)

	var res3 *Res

	res4 := &Res{
		Status: "failed",
	}

	fmt.Println(res1, res2, res3, res4)
	fmt.Printf("%+v %+v %+v %+v", res1, res2, res3, res4)
}

/* result:
&{ } &{ } <nil> &{failed }
&{Status: Msg:} &{Status: Msg:} <nil> &{Status:failed Msg:}
*/
