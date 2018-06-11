package main

import (
	"math"
	"fmt"
)

type Vertex2 struct {
	X, Y float64
}

func Abs(v Vertex2) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex2{3, 4}
	fmt.Println(Abs(v))
}
