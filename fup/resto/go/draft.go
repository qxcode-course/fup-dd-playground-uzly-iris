package main

import "fmt"

func main() {
	var b, c, q, r int
	fmt.Scan(&b)
	fmt.Scan(&c)
	q = b / c
	r = b % c
	fmt.Println(q, r)
}
