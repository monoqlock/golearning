package funcs

import "fmt"

func ExecPointer() {

	var i int
	p := &i
	fmt.Printf("p=%T\n", p)
	fmt.Printf("*p=%T\n", *p)
	i = 5
	fmt.Printf("p=%v\n", p)
	fmt.Printf("*p=%v\n", *p)

	p2 := &[3]string{"Apple", "Banana", "Orange"}
	for i2, v2 := range p2 {
		fmt.Printf("i2=%d, v2=%v\n", i2, v2)
	}

}

func inc(i *int) {
	*i++
}

func ExecInc() {
	i := 0
	inc(&i)
	inc(&i)
	inc(&i)
	fmt.Println(i)

}

func pow(a *[3]int) {

	i := 0
	for i < len(*a) {
		(*a)[i] = (*a)[i] * (*a)[i]
		i++
	}
}

func ExecPow() {

	a := [3]int{1, 2, 3}
	pow(&a)
	fmt.Println(a)
}
