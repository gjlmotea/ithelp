package main

import "fmt"

func main()  {
	s1 := "I"
	s2 := "am"
	s3 := "string"

	str1 := fmt.Sprintln(s1, s2, s3)
	fmt.Println(str1)

	str2 := fmt.Sprint(s1, s2, s3)
	fmt.Println(str2)
}
/* result:
I am string

Iamstring

*/