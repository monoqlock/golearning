package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {

	var res [][]uint8

	for i := 0; i < dx; i++ {
		var a []uint8
		for j := 0; j < dy; j++ {
			a = append(a, uint8((i+j)/2))
		}
		res = append(res, a)
	}

	return res
}

func main() {
	pic.Show(Pic)
}
