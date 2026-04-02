package main

import "fmt"

func main() {
	var n, d, a int
	fmt.Scan(&n)
	fmt.Scan(&d)
	fmt.Scan(&a)

	if d < a {
		c := (n - a) + d
		fmt.Println(c)
	} else if d > a {
		c := d - a
		fmt.Println(c)
	} else {
		fmt.Println(0)
	}
}
