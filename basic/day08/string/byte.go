package main

import "fmt"

func main()  {
	var b = []byte("0")
	fmt.Println(b)

	b = []byte("爆肝工程師的異世界安魂曲")
	fmt.Println(b)
}

/* result:
[48]
[231 136 134 232 130 157 229 183 165 231 168 139 229 184 171 231 154 132 231 149 176 228 184 150 231 149 140 229 174 137 233 173 130 230 155 178]
 */