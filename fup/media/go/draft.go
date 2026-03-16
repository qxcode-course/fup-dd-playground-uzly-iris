package main

import "fmt"

func main() {
	var n1, n2, m float64
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	m = (n1 + n2) / 2
	fmt.Printf("%.1f\n", m)
}
