package main

import (
	"fmt"
)

func main() {
	var s float64
	fmt.Scan(&s)

	if s <= 1000 {
		s = s + (s / 5)
		fmt.Printf("%.2f\n", s)
	} else if s <= 1500 {
		s = s + (s / 20) * 3
		fmt.Printf("%.2f\n", s)
	} else if s <= 2000 {
		s = s + (s / 10)
		fmt.Printf("%.2f\n", s)
	} else {
		s = s + (s / 20)
		fmt.Printf("%.2f\n", s)
	}
}
