package main

import "fmt"

func main() {
	var n1, n2, n3, m1, m2 float64
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	fmt.Scan(&n3)

	m1 = (n1 + n2) / 2

	if m1 >= 7 {
		fmt.Println("aprovado")
	} else if m1 < 4 {
		fmt.Println("reprovado")
	} else {
		m2 = (m1 + n3) / 2

		if m2 >= 5 {
			fmt.Println("aprovado na final")
		} else {
			fmt.Println("reprovado na final")
		}
	}
}
