package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 float64
	fmt.Scan(&x1, &y1, &x2, &y2)

	X := math.Pow((x2 - x1), 2)
	Y := math.Pow((y2 - y1), 2)
	d := math.Sqrt(X + Y)
	fmt.Printf("%.2f\n", d)
}
