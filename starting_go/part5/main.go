package main

import (
	"./funcs"
	"fmt"
)

func main() {
	fmt.Println("----pointer")
	funcs.ExecPointer()
	funcs.ExecInc()
	funcs.ExecPow()

	fmt.Println("------struct")
	funcs.ExecTypes()
	funcs.ExecSum()
	funcs.ExecStructs()
	funcs.ExecTags()
}
