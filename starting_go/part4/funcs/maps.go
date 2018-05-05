package funcs

import (
	"fmt"
)

func Maps() {
	m := make(map[int]string)
	m[1] = "Apple"
	m[10] = "Orange"

	fmt.Println(m)

	m2 := map[int]string{1: "Taro", 2: "Hanako", 3: "Jiro"}
	fmt.Println(m2)
	s, ok := m2[2]
	s2, ok2 := m2[8]
	fmt.Printf("s=%v, ok=%v\n", s, ok)
	fmt.Printf("s2=%v, ok2=%v\n", s2, ok2)

	if _, ok := m2[3]; ok {
		fmt.Println("m[3] OK")
	}

	for k, v := range m2 {
		fmt.Printf("k=%v => v=%v\n", k, v)
	}

	delete(m2, 2)
	fmt.Println(m2)
}
