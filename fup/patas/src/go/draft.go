package main

import (
    "fmt"
    "math"
)

func main() {
	var qtd int
    var cebolinha, CB float64
	fmt.Scan(&CB, &cebolinha, &qtd)
	animais := make([]string, qtd)

	for i := range animais {
		fmt.Scan(&animais[i])
	}
    pernas(animais, cebolinha, CB)
}

func pernas(animais[] string, cebolinha float64, CB float64) {
    var cont float64
    for _, i := range animais {
        switch i {
        case "v":
            cont += 4
        case "g":
            cont += 2
        case "c":
            cont +=4 
        }
    }
    fmt.Println(cont)
    if math.Abs(cont - CB) > math.Abs(cont - cebolinha) {
        fmt.Println("Cebolinha")
    } else if math.Abs(cont - CB) == math.Abs(cont - cebolinha) {
        fmt.Println("empate")
    } else {
        fmt.Println("Chico Bento")
    }
    
}