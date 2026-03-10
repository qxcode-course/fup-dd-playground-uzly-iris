package main

import "fmt"

func main() {
	var b, c int
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Print(b/c, b%c)
}
