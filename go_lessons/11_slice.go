package main

import "fmt"

// 11 スライスを使ってみよう
func main() {

	a := [5]int{2, 5, 7, 3}
	s := a[2:4]

	fmt.Println("a :", a) // a : [2 5 7 3 0]
	fmt.Println("s :", s) // s : [7 3]

	s[1] = 12
	fmt.Println(s) // [7 12]

	fmt.Println(len(s)) // 2
	fmt.Println(cap(s)) // 3 切り出しうる個数

}
