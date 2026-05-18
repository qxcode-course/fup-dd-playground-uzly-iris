package main

import "fmt"

func main() {
	var qtd float64
	fmt.Scan(&qtd)
	cals := make([]float64, int(qtd))

	for i := range cals {
		fmt.Scan(&cals[i])
	}
    m := media_cals(cals, qtd)
    fmt.Printf("%.1f\n", m)
}

func media_cals(cals[] float64, qtd float64) float64 {
    var i, x float64
    for _, i = range cals {
        x += i
    }
    m := x / qtd
    return m
}
