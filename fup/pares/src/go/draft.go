package main

import "fmt"

func main() {
	var a, b, acc int
	fmt.Scan(&a)
	fmt.Scan(&b)

	if a > b {
		fmt.Println("invalido")
		return
	}

	acc = soma_par(a, b)
	fmt.Println(acc)
}

func soma_par(a, b int) int {
    acc := 0
    for x := a ; x <= b ; x++ {
        if x%2 == 0 {
            acc += x
            continue
        }
    }
    return acc
}
