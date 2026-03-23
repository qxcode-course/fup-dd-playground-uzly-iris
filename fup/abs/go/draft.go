package main

import (
	"fmt"
	"math"
)

func main() {
	var n1, n2 float64
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	x := int(math.Abs(n1 - n2))
	fmt.Println(x)

}
