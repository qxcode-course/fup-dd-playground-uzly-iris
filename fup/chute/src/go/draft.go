package main

import (
	"fmt"
	"math"
)

func main() {
	var p, r1, r2 float64
	fmt.Scan(&p)
	fmt.Scan(&r1)
	fmt.Scan(&r2)

	c1 := int(math.Abs(p - r1))
	c2 := int(math.Abs(p - r2))

	if c1 == c2 {
		fmt.Println("empate")
	} else if c1 < c2 {
		fmt.Println("primeiro")
	} else {
		fmt.Println("segundo")
	}
}
