package main

import "fmt"

func main() {
	var a, b, smt int
	fmt.Scan(&a)
	fmt.Scan(&b)

	if a > b {
		fmt.Println("invalido")
		return
	}

	smt = dividir_por_2_e_3(a, b)
	fmt.Println(smt)
}

func dividir_por_2_e_3(a, b int) int {
	smt := 0
	for x := a; x <= b; x++ {
		if x%2 == 0 && x%3 == 0 {
			smt += x
			continue
		}
	}
	return smt
}
