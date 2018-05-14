package main

import "fmt"

// 07 ポインタを使ってみよう
func main() {

	a := 3
	var pa *int
	pa = &a // aのアドレス
	// *pa : paの領域にあるデータの値
	fmt.Println(pa)
	fmt.Println(*pa)

}
