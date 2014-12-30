package main

import "fmt"

type user struct {
	name string
	score int
}

func main() {
	u := new(user)
	fmt.Println(u) // &{ 0}

	u.name = "morikawa"
	u.score = 20
	fmt.Println(u) // &{morikawa 20}

	u2 := user{"taguchi", 100}
	fmt.Println(u2) // {taguchi 100}
}