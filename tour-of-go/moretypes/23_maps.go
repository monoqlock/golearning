package main


import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {

	res := make(map[string]int)
	for _, v := range strings.Fields(s) {
		if _, ok := res[v]; ok {
			res[v] = res[v] + 1
		} else {
			res[v] = 1
		}
	}

	return res
}

func main() {
	wc.Test(WordCount)
}