package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Print("[ ")
	for ; a < b; a++ {
		fmt.Print(a, " ")
	}
	fmt.Print("]\n")
}
