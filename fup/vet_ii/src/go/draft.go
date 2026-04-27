package main

import "fmt"

func main() {
	var qtd int
	fmt.Scan(&qtd)
	numeros := make([]int, qtd)

	for i := range numeros {
		fmt.Scan(&numeros[i])
	}

    fmt.Print("[ ")
	if qtd == 0 {
		fmt.Print("")
	} else {
        for _, valor := range numeros {
                fmt.Printf("%v ", valor)
        }
    }
    fmt.Print("]\n")
}
