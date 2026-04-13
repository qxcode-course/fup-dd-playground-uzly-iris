package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	A := a
	B := b
	fmt.Print("[ ")
	for i := a; i <= b; i++ {
		fmt.Print(A, " ")
		A += 1
		fmt.Print(B, " ")
		B -= 1
	}
	fmt.Print("]\n")
}
