package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	b2 := math.Pow(b, 2)
	d = b2 - 4*a*c

	if d > 0 {
		rd := math.Sqrt(d)
		bskp := ((-1 * b) + rd) / (2 * a)
		bskn := ((-1 * b) - rd) / (2 * a)

		fmt.Printf("%.2f\n", bskp)
		fmt.Printf("%.2f\n", bskn)
	} else if d == 0 {
		rd := math.Sqrt(d)
		bsk := ((-1 * b) + rd) / (2 * a)

		fmt.Printf("%.2f\n", bsk)
	} else {
		fmt.Println("nao ha raiz real")
	}
}
