package main

import "fmt"

// 09 関数を使ってみよう(2)

func swap(a, b int) (int, int) {
	return b, a
}

func main() {

	fmt.Println(swap(1, 3))

	f := func(a, b int) (int, int) {
		return b, a
	}
	fmt.Println(f(10, 20))

	func (msg string) {
		fmt.Println(msg)
	} ("morikawa")
}