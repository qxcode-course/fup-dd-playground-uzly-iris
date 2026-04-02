package main

import (
	"fmt"
	"math"
)

func main() {
	var o string
	var n float64
	fmt.Scan(&o)
	fmt.Scan(&n)

	switch o {
	case "r":
		n = math.Round(n)
		fmt.Println(n)
	case "f":
		n = math.Floor(n)
		fmt.Println(n)
	default:
		n = math.Ceil(n)
		fmt.Println(n)
	}
}
