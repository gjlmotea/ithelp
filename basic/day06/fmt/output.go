package main

import "fmt"

func main()  {
	a := 10
	fmt.Printf("a: %d\n", a)
	fmt.Println("a: ", a)

	s1 := "I"
	s2 := "am"
	s3 := "string"
	fmt.Printf("%s%s      %s\n",s1, s2, s3)
	fmt.Println(s1 + s2 + s3)
	fmt.Println(s1, s2, s3)

	fmt.Println("========")

	fmt.Print(s1 + s2 + s3)
	fmt.Print(s1, s2, s3)
}
/* result:
a: 10
a:  10
Iam      string
Iamstring
I am string
*/