package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["taguchi"] = 200
	m["morikawa"] = 100

	fmt.Println(m) // map[taguchi:200 morikawa:100]

	a := map[string]int{"taguchi": 200, "fkoji": 300}
	fmt.Println(a)
	fmt.Println(len(a)) // 2

	delete(a, "fkoji")
	fmt.Println(a) // map[taguchi:200]

	v, ok := a["fkoji"]
	v2, ok2 := a["taguchi"]
	fmt.Println(v)   // 0
	fmt.Println(ok)  // false
	fmt.Println(v2)  // 200
	fmt.Println(ok2) // true
}
