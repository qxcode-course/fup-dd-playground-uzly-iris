package main

import "fmt"

func main() {
	var a, b, c, h, l int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&h)
	fmt.Scan(&l)

	j := h * l

	if j > a*b || j > a*c || j > b*c {
		fmt.Println("S")
	} else {
		fmt.Println("N")
	}
}
