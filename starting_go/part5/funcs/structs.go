package funcs

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type (
	IntPair     [2]int
	AreaMap     map[string][2]float64
	IntsChannel chan int
)

func (ip IntPair) Plus(ip2 IntPair) int {
	return ip[1] + ip2[1]
}

func ExecTypes() {
	pair := IntPair{1, 2}
	area := AreaMap{"Tokyo": {35.7, 139.7}}
	ichan := make(IntsChannel)
	fmt.Println(pair)
	fmt.Println(area)
	fmt.Println(ichan)

}

type CallBack func(i int) int

func sum(a []int, back CallBack) []int {
	for i, v := range a {
		a[i] = back(v)
	}
	return a
}

func ExecSum() {
	a := []int{1, 2, 3}
	cb := func(i int) int { return i + i }
	b := sum(a, cb)
	fmt.Println(b)
}

type Point struct {
	X int
	Y int
}

func (point *Point) Render() {
	fmt.Printf("<%d, %d>\n", point.X, point.Y)
}

type Feed struct {
	Name   string
	Amount int
}

type Animal struct {
	Name string
	Feed Feed
}

func NewAnimal(name string, feed Feed) *Animal {
	a := new(Animal)
	a.Name = name
	a.Feed = feed
	return a
}

func (point *Point) Set(x int, y int) {
	point.X = x
	point.Y = y
}

type Points []*Point

func (points Points) ToString() string {
	str := ""
	fmt.Println(points)
	for _, v := range points {

		if str != "" {
			str += ","
		}

		if v == nil {
			str += "<nil>"
		} else {
			str += fmt.Sprintf("[%d, %d]", v.X, v.Y)
		}
	}

	return str
}

func ExecStructs() {

	p := Point{X: 1, Y: 2}
	p.X = 3
	p.Y = 4
	fmt.Println(p)

	a := Animal{Name: "monkey", Feed: Feed{Name: "Banana", Amount: 3}}
	a.Feed.Amount = 4
	fmt.Println(a)

	p2 := new(Point)
	p2.X = 5
	p2.Y = 6
	fmt.Println(p2) // &{5 6}

	p2.Render()

	ip := IntPair{1, 2}
	ip2 := IntPair{3, 4}
	ip3 := ip.Plus(ip2)
	fmt.Printf("ip3=%d\n", ip3)

	fmt.Println(NewAnimal("Monkey", Feed{Name: "Banana", Amount: 3}))

	p3 := new(Point)
	p3.Set(1, 2)
	fmt.Print("p3.X=%d\n", p3.X)
	fmt.Print("p3.Y=%d\n", p3.Y)

	ps := make([]Point, 5)
	for _, v := range ps {
		fmt.Println(v)
	}

	ps2 := Points{}
	ps2 = append(ps2, &Point{1, 2})
	ps2 = append(ps2, nil)
	ps2 = append(ps2, &Point{3, 4})
	fmt.Println(ps2.ToString())

	m1 := map[Point]string{
		{1, 2}: "test",
		{3, 4}: "test2",
	}
	fmt.Println(m1)
}

type User struct {
	Id   int    `json:"user_id"`
	Name string `json:"user_name"`
	Age  uint   `json:"user_age"`
}

func ExecTags() {

	u := User{123, "Taro", 20}

	t := reflect.TypeOf(u)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println("Tag: ", f.Name, f.Tag)
	}

	bs, _ := json.Marshal(u)
	fmt.Println("JSON: ", string(bs))
}
