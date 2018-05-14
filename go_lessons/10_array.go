package main

import "fmt"

// 10 配列を使ってみよう
func main() {

	var a [5]int
	a[2] = 3
	a[3] = 10
	fmt.Println(a) // [0 0 3 10 0]
	fmt.Println(a[2])

	b := [3]int{1, 3, 5}
	fmt.Println(b)

	c := [...]int{2, 4, 6, 8}
	fmt.Println(c)
	fmt.Println(len(c))

}
