package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	if a < b {
		fmt.Print("[ ")
		for ; a < b; a++ {
			fmt.Print(a, " ")
		}
		fmt.Print("]\n")

	} else {
		fmt.Print("[ ")
		for ; b < a; a-- {
			fmt.Print(a, " ")
		}
		fmt.Print("]\n")
	}
}
