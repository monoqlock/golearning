package main

import "fmt"

func main() {
	var i interface{}
	describe14(i) // (<nil>, <nil>)

	i = 42
	describe14(i) // (42, int)

	i = "hello"
	describe14(i) // (hello, string)
}

func describe14(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
