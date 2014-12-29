package main

import "fmt"

// 08 関数を使ってみよう

func hi(name string) string {
	msg := "hi!" + name
	fmt.Println(msg)
	// fmt.Println("hi! " + name)
	return msg
}

func hi2(name string) (msg string) {
	msg = "hello! " + name
	fmt.Println(msg)
	return
}

func main() {
	hi("morikawa")
	hi2("kaori")
}