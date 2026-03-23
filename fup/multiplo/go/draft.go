package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	if x%7 == 0 {
		fmt.Println("SIM")
	} else {
		fmt.Println("NAO")
	}
}
