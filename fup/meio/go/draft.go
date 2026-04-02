package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	if a < b && b < c || c < b && b < a {
		fmt.Println(b)
	} else if b < a && a < c || c < a && a < b {
		fmt.Println(a)
	} else {
		fmt.Println(c)
	}
}
