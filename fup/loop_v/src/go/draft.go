package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Print("[ ")
	for {
		if a%2 != 0 {
			fmt.Print(a, " ")
			a = a + 1
			continue
		} else if a == b {
			break
		} else {
			a = a + 1
			continue
		}
	}
	fmt.Print("]\n")
}
