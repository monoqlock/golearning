package funcs

import (
	"fmt"
)

type MyError struct {
	Message   string
	ErrorCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{"エラーが発生しました", 1234}
}

type Stringify interface {
	ToString() string
}

type Person struct {
	Name string
	Age  uint
}

func (p *Person) ToString() string {
	return fmt.Sprintf("%s(%d)", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("[%s] %s", c.Number, c.Model)
}

func Println(s Stringify) {
	fmt.Println(s.ToString())
}

type T struct {
	Id   int
	Name string
}

func (t *T) String() string {
	return fmt.Sprintf("<<%d, %s>>", t.Id, t.Name)
}

func (t *T) GetId() int {
	return t.Id
}

type I0 interface {
	Method01() int
}

type I2 interface {
	I0
	Method02() int
}

type I3 interface {
	I2
	Method03() int
}

func showId(t interface{ GetId() int }) {
	fmt.Println("ID=", t.GetId())
}

func ExecInterface() {
	err := RaiseError()
	fmt.Println(err.Error())

	vs := []Stringify{
		&Person{"Tom", 21},
		&Car{"1234", "Px512"},
	}
	for _, v := range vs {
		fmt.Println(v.ToString())
	}

	Println(&Person{"Jhon", 18})

	t := &T{123, "Alice"}
	fmt.Println(t)

	showId(t)
}
