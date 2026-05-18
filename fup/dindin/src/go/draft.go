package main

import "fmt"

func main() {
	var qtd, contSC, contSL, contTM, contTT int
	fmt.Scan(&qtd)
    sabor := make([]string, qtd)
	turno := make([]string, qtd)

	for i := 0 ; i < qtd ; i++ {
        fmt.Scan(&sabor[i])
		fmt.Scan(&turno[i])
	}

    for _, i := range sabor {
        if i == "c" {
            contSC += 1
        } else {
            contSL += 1
        }
    }

    for _, i := range turno {
        if i == "m" {
            contTM += 1
        } else {
            contTT += 1
        }
    }

    if contSC > contSL {
        fmt.Println("c")
    } else if contSC == contSL {
        fmt.Println("empate")
    } else {
        fmt.Println("l")
    }

    if contTM > contTT {
        fmt.Println("t")
    } else if contTM == contTT {
        fmt.Println("empate")
    } else {
        fmt.Println("m")
    }
}