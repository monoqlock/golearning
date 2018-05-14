package main

import "fmt"

func main() {

	signal := "ww"
	switch signal {
	case "red":
		fmt.Println("Stop")
	case "blue", "green":
		fmt.Println("Go")
	case "yellow":
		fmt.Println("Caution")
	default:
		fmt.Println("wrong signal")
	}

	score := 80
	switch {
	case score > 80:
		fmt.Println("Great!")
	case score > 69:
		fmt.Println("Good")
	default:
		fmt.Println("so so..")
	}
}
