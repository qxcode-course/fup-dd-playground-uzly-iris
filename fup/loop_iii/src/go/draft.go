package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Print("[ ")
	for ; b < a; a-- {
		fmt.Print(a, " ")
	}
	fmt.Print("]\n")
}