package main

import (
	"./animals"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!!")
	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.RabbitFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(AppName())
}
