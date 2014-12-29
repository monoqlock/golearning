package main

import "fmt"

// 05 定数を使ってみよう
func main() {

	const name = "morikawa"
	//name = "fkoji"
	fmt.Println(name)

	const (
		sun = iota
		mon
		tue
	)
	fmt.Println(sun, mon, tue)
}
