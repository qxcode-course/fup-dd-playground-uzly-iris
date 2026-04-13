package main

import "fmt"

func main() {
	var a, b, acc int
	fmt.Scan(&a, &b)

	if a <= b {
		for i := a; i <= b; i++ {
			if i%2 == 0 {
				acc += i
			}
		}
		fmt.Println(acc)
	} else {
		fmt.Println("invalido")
	}
}
