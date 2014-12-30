package main

import "fmt"

func show(t interface{}) {
	// 型アサーション
	// _, ok := t.(japanese)
	// if ok {
	// 	fmt.Println("私は日本人です")
	// } else {
	// 	fmt.Println("I am not Japanese")
	// }

	// 型スイッチ
	switch t.(type) {
	case japanese:
		fmt.Println("私は日本人です")
	case american:
		fmt.Println("I am American")
	}
}

type greeter interface {
	greet()
}

type japanese struct {}
type american struct {}

func (j japanese) greet() {
	fmt.Println("こんにちは")
}

func (a american) greet() {
	fmt.Println("Hello")
}

func main() {
	greeters := []greeter{japanese{}, american{}}
	for _, greeter := range greeters {
		show(greeter)
	}
}