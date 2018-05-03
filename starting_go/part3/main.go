package main

import (
	"fmt"
	"math"
)

func init() {
	fmt.Println("init!!")
}

func main() {

	fmt.Printf("five=%d\n", 5)
	fmt.Printf("10進数=%d, 2進数=%b, 8進数=%o, 16進数=%x\n", 17, 17, 17, 17)
	fmt.Printf("数値=%v, 文字列=%v, 配列=%v\n", 5, "Golang", [...]int{1, 2, 3})
	fmt.Printf("数値=%#v, 文字列=%#v, 配列=%#v\n", 5, "Golang", [...]int{1, 2, 3})
	fmt.Printf("数値=%T, 文字列=%T, 配列=%T\n", 5, "Golang", [...]int{1, 2, 3})

	ui_1 := uint32(400000000)
	ui_2 := uint32(4000000000)
	sum := ui_1 + ui_2
	fmt.Printf("%d + %d = %d\n", ui_1, ui_2, sum) // 400000000 + 4000000000 = 105032704

	fmt.Printf("uint32 max value = %d\n", math.MaxUint32)

	var x interface{}
	fmt.Printf("%#v\n", x)

	fmt.Println(plus(1, 3))
	hello()
	q, r := div(7, 2)
	fmt.Printf("商=%d, 余り=%d\n", q, r)

	x, y := doSomething()
	fmt.Printf("x=%d, y=%d\n", x, y)

	fmt.Printf("%T\n", func(x, y int) int { return x + y })
	fmt.Printf("%#v\n", func(x, y int) int { return x + y })
	fmt.Printf("%#v\n", func(x, y int) int { return x + y }(3, 2))

	f := later()
	fmt.Println(f("Golang"))    // ""
	fmt.Println(f("is"))        // Golang
	fmt.Println(f("awesome!!")) // is

	ints := integers()
	fmt.Println(ints()) // 1
	fmt.Println(ints()) // 2

	fmt.Println("---loop")
	loop()

	fmt.Println("--range")
	ranges()

	sw()

	assertion()

	label()

	labelLoop()

	runDefer()

	go subLoop()
	for i := 0; i < 10; {
		fmt.Println("main loop")
		i++
	}

	printRuntime()
}

func plus(x, y int) int {
	return x + y
}

func hello() {
	fmt.Println("Hello!!")
	return
}

func div(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}

func doSomething() (x, y int) {
	y = 5
	return // x = 0, y = 5
}

func later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func integers() func() int {
	i := 0
	return func() int {
		i += 1
		return i 
	}
}



