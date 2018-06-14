package main

import "fmt"

type I13 interface {
	M13()
}

func main() {
	var i I13
	describe13(i)
	i.M13()
}

func describe13(i I13) {
	fmt.Printf("(%v, %T)\n", i, i) // error
}