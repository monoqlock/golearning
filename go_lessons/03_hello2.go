// 変数

package main

import "fmt"

// 03 変数を使ってみよう
func main() {

	//var msg string
	//msg = "Hello world"
	// var msg = "Hello World!!"
	msg := "hello world !!"
	fmt.Println(msg)

	// var a, b int
	// a, b = 1, 3
	a, b := 2, 4
	fmt.Println(a, b)

	var (
		c int
		d string
	)
	c = 20
	d = "hoge"
	fmt.Println(c, d)

}
