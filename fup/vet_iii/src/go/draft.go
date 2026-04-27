package main

import "fmt"

func main() {
	var qtd int
	fmt.Scan(&qtd)
	numeros := make([]int, qtd)

	for i := range numeros {
		fmt.Scan(&numeros[i])
	}

    fmt.Print("[")
	if qtd == 0 {
		fmt.Print("")
	} else {
        for i, valor := range numeros {
            if i == qtd - 1 {
                fmt.Printf("%v", valor)
            } else {
                fmt.Printf("%v, ", valor)
            }
        }
    }
    fmt.Print("]\n")
}
