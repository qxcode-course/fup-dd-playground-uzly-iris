package main

import "fmt"

func main() {
	// numeros := []int{6, 0, 8}
	// fmt.Println(numeros)
	// fmt.Println(numeros[0])

	// var qtd int
	// fmt.Scan(&qtd)
	// var idades []int = make([]int, qtd)
	// fmt.Println(idades)

	spcs := []string{"yamoki", "leo", "sani", "miguel"}
	fmt.Print("[\n")
	for pos, valor := range spcs {
		fmt.Printf("%v %v\n", pos, valor)
	}
	fmt.Print("]\n")
}
