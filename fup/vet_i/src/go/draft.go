package main

import "fmt"

func main() {
	var qtd int
	fmt.Scan(&qtd)
	numeros := make([]int, qtd)

	for i := range numeros {
		fmt.Scan(&numeros[i])
	}

	if qtd == 0 {
		fmt.Println("")
	}
	for _, valor := range numeros {
		if qtd != 0 {
			fmt.Printf("%v\n", valor)
		}
	}
}
