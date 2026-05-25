package main

import (
	"fmt"
)

func main() {
    var qtd int
    fmt.Scan(&qtd)
    
    balatro := make([]string, qtd)

    for i := 0 ; i < qtd; i++ {
        fmt.Scan(&balatro[i])

        switch balatro[i] {
        case "1":
            balatro[i] = "A"
        case "11":
            balatro[i] = "J"
        case "12":
            balatro[i] = "Q"
        case "13":
            balatro[i] = "K"
        }
    }

    fmt.Print("[")
    if qtd == 0 {
        fmt.Print("")
    } else {
        for i := 0 ; i < qtd ; i++ {
            if i == qtd - 1 {
                fmt.Print(balatro[i])
                break
            }
            fmt.Print(balatro[i], ", ")
        }
    }
    fmt.Print("]\n")
}