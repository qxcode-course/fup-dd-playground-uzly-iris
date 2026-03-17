package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, d, qx, qy float64
	fmt.Scan(&x1, &y1)
	fmt.Scan(&x2, &y2)
	qx = math.Pow((x2 - x1), 2)
	qy = math.Pow((y2 - y1), 2)
	d = math.Sqrt(qx + qy)
	fmt.Printf("%.2f\n", d)
}
