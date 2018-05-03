package main

import (
	"fmt"
	"runtime"
)

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 5 {
			break
		}
	}
}

func ranges() {
	fruits := [3]string{"Apple", "Orange", "Banana"}
	for i, s := range fruits {
		fmt.Printf("%d: %s\n", i, s)
	}
}

func sw() {
	num := 5
	switch num {
		case 1:
		 fmt.Println("One")
		case 5:
		fmt.Println("Five")
		default:
		fmt.Println("Default")
	}
}

func assertion() {
	var as interface{} = "assertion"
	switch v := as.(type) {
		case int, uint:
			fmt.Printf("int: value=%d\n", v)
		case string:
			fmt.Printf("string: value=%s\n", v)
		case bool:
			fmt.Printf("bool: value=%v\n", v)
		default:
			fmt.Println("No type matche")
	}
}

func label() {
	fmt.Println("A")
	goto L
	fmt.Println("B")
	L:
	fmt.Println("C")
}

func labelLoop() {
	LABEL:
	for {
		for {
			fmt.Println("start in label")
			break LABEL
		}
		fmt.Println("Not Pass")
	}
}

func runDefer() {
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	fmt.Println("done")
}

func subLoop() {
	for j := 0; j < 5;{
		fmt.Printf("subLoop: %d\n", j)
		j++
	}
}


func printRuntime() {
	go fmt.Println("goroutine")
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
	fmt.Printf("Version: %s\n", runtime.Version())
}
