package main

import "fmt"

func main() {
	var ct, a, b, c int
	fmt.Scan(&ct)
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	r := ((a + b + c) % ct)

	if r == 0 {
		m := ((a + b + c) / ct)
		fmt.Println(m)
	} else {
		m := ((a + b + c) / ct) + 1
		fmt.Println(m)
	}
}
