package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, d, x, y, qx, qy float64
	fmt.Scan(&x1)
	fmt.Scan(&y1)
	fmt.Scan(&x2)
	fmt.Scan(&y2)
	x = x2 - x1
	y = y2 - y1
	qx = math.Pow(x, 2)
	qy = math.Pow(y, 2)
	d = math.Sqrt(qx + qy)
	fmt.Printf("%.2f\n", d)
}
