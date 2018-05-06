package funcs

import "fmt"

type (
	IntPair [2]int
	AreaMap map[string][2]float64
	IntsChannel chan int
)

func (ip IntPair) Plus(ip2 IntPair) int {
	return ip[1] + ip2[1]
}

func ExecTypes()  {
	pair := IntPair{1,2}
	area := AreaMap{"Tokyo" : {35.7, 139.7}}
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

func ExecSum()  {
	a := []int{1,2,3}
	cb := func(i int) int { return i + i}
	b := sum(a,cb)
	fmt.Println(b)
}

type Point struct{
	X int
	Y int
}

func (point *Point) Render()  {
	fmt.Printf("<%d, %d>\n", point.X, point.Y)
}

type Feed struct {
	Name string
	Amount int
}

type Animal struct {
	Name string
	Feed Feed
}


func ExecStructs()  {

	p := Point{X: 1, Y:2}
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

	ip := IntPair{1,2}
	ip2 := IntPair{3,4}
	ip3 := ip.Plus(ip2)
	fmt.Printf("ip3=%d\n", ip3)
}

