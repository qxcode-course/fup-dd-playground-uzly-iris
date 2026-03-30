package main

import (
	"fmt"
)

func main() {
	var n, x, y, s int
	var c string
	fmt.Scan(&n)
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Scan(&c)
	fmt.Scan(&s)

	if c == "U" || c == "D" {
		if c == "U" {
			y = y - (s % n)
			y = y % n

			fmt.Println(x, y)
		} else {
			y := y + s
			y = y % n

			fmt.Println(x, y)
		}
	} else {
		if c == "L" {
			x = x - (s % n)

			fmt.Println(x, y)
		} else {
			x = x + s
			x = x % n

			fmt.Println(x, y)
		}

	}
}
