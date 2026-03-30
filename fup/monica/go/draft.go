package main

import "fmt"

func main() {
	var m, a, b, c int
	fmt.Scan(&m)
	fmt.Scan(&a)
	fmt.Scan(&b)

	c = m - (a + b)

	fmt.Println(c)
}
