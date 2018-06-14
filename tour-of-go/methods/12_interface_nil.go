package main

import "fmt"

type I12 interface {
	M12()
}

type T12 struct {
	S12 string
}

func (t *T12) M12() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S12)
}

func main() {
	var i I12

	var t *T12
	i = t
	describe12(i)
	i.M12()

	i = &T12{"hello"}
	describe12(i)
	i.M12()
}

func describe12(i I12) {
	fmt.Printf("(%v, %T)\n", i, i)
}