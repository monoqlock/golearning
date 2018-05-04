package main

import (
	"fmt"
)

func slice() {
	s := make([]int, 10)
	fmt.Println(s)
	a := [10]int{}
	fmt.Println(a)

	s[1] = 3
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	s2 := make([]int, 2, 5)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))

	s3 := []int{1,3,5,7,9}
	s4 := s3[2:4]
	fmt.Println(s4)

	s5 := "ABCDE"[2:4]
	fmt.Println(s5)

	// append
	s2 = append(s2, 7)
	fmt.Println(s2)
	s2 = append(s2, 1, 2)
	fmt.Println(s2)
	s6 := append(s2, s3...)
	fmt.Println(s6)

	s7 := make([]int, 0, 0)
	fmt.Printf("A) s7=%v, len=%d, cap=%d\n", s7, len(s7), cap(s7))
	s7 = append(s7, 1)
	fmt.Printf("A) s7=%v, len=%d, cap=%d\n", s7, len(s7), cap(s7))
	s7 = append(s7, []int{2,3,4}...)
	fmt.Printf("B) s7=%v, len=%d, cap=%d\n", s7, len(s7), cap(s7))
	s7 = append(s7, 5)
	fmt.Printf("C) s7=%v, len=%d, cap=%d\n", s7, len(s7), cap(s7))

	// copy
	s8 := []int{1,2,3,4,5,6}
	s9 := []int{4,5,6}
	copy(s8, s9)
	fmt.Println(s8)
}

func slice2() {
	s1 := []int{1,2,3,4,5,6,7,8}
	s2 := s1[2:4:6]
	fmt.Printf("s2=%v, len=%d, cap=%d\n", s2, len(s2), cap(s2))

	s3 := []string{"Apple", "Banana", "Orange"}
	for i, v := range s3 {
		fmt.Printf("[%d] = %s\n", i, v)
	}

}

func pow(a []int) {
	for i, v := range a {
		a[i] = v * v
	}
}

func execPow() {
	s := []int{1, 2, 3}
	pow(s)
	fmt.Println(s)
}
