package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		if i > 4 {
			break
		}
		fmt.Println(i)
	}

	j := 0
	for j < 4 {
		fmt.Println(j)
		j++
	}

	k := 0
	for {
		fmt.Println(k)
		k++
		if k > 5 {
			break
		}
	}

}
