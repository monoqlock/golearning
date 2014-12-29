package main

import "fmt"

// 12 make()､append()､copy()を使おう
func main() {

	// make(型, len, cap)
	// s := make([]int, 3, 5)
	s := make([]int, 3)
	fmt.Println(s)

	// 追加
	a := []int{1, 5, 12}
	a = append(a, 2, 4, 6)
	fmt.Println("a :", a) // a : [1 5 12 2 4 6]

	// copy
	b := make([]int, len(a))
	c := copy(b, a)
	fmt.Println("b :", b) // b : [1 5 12 2 4 6]
	fmt.Println("c :", c) // 6 コピーした要素の数を返す
}