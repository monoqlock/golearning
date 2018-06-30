package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func walk(t *tree.Tree, ch chan int) {

	if t == nil {
		return
	}

	if t.Left != nil {
		walk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		walk(t.Right, ch)
	}
}

func Same(t1 *tree.Tree, t2 *tree.Tree) bool {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	s1 := make([]int, 0, 10)
	s2 := make([]int, 0, 10)

	for c1 := range ch1 {
		s1 = append(s1, c1)
	}
	for c2 := range ch2 {
		s2 = append(s2, c2)
	}

	for i := range s1 {
		if len(s2) < i || s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func main() {

	t := tree.New(1)
	fmt.Println(t)

	ch := make(chan int)
	go Walk(t, ch)

	for i := range ch {
		fmt.Println(i)
	}


	t1 := tree.New(1)
	t2 := tree.New(2)
	fmt.Println(Same(t, t1)) // true
	fmt.Println(Same(t1, t2)) // false
}
