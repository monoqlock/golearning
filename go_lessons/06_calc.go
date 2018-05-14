package main

import "fmt"

// 06 基本的な演算をしてみよう
func main() {

	var a int
	a = 10
	// a++
	// a = a % 3
	a = a / 3

	fmt.Println(a)

	var s string
	s = "hello" + "world"
	fmt.Println(s)

	t := true
	f := false
	fmt.Println("t && f :", t && f)
	fmt.Println("t || f :", t || f)
	fmt.Println("!t :", !t)
}
