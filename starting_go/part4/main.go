package main

import (
	"fmt"
	"./funcs"
)

func main() {
	fmt.Println("Part4")

	fmt.Println("-------- slice")
	funcs.Slice()
	funcs.Slice2()
	funcs.ExecPow()

	fmt.Println("-------- map")
	funcs.Maps()

	fmt.Println("---------channels")
	funcs.Channels()
	funcs.Select()
	funcs.SelectMix()
}
