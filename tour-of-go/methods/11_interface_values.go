package main

import (
	"fmt"
	"math"
)

type I11 interface {
	M11()
}

type T11 struct {
	S string
}

func (t *T11) M11() {
	fmt.Println(t.S)
}

type F11 float64

func (f F11) M11() {
	fmt.Println(f)
}

func main() {
	var i I11

	i = &T11{"Hello"}
	describe(i)
	i.M11()

	i = F11(math.Pi)
	describe(i)
	i.M11()
}

func describe(i I11) {
	fmt.Printf("(%v, %T)\n", i, i)
}
