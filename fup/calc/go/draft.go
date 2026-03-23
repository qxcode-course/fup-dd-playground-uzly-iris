package main

import "fmt"

func main() {
	var n1, n2, r int
	var o string
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	fmt.Scan(o)

	switch o {
    case "x":
		r = n1 + n2
		fmt.Println(r)
	case "-":
		r = n1 - n2
		fmt.Println(r)
	case "*":
		r = n1 * n2
		fmt.Println(r)
	default:
		r = n1 / n2
		fmt.Println(r)
	}
}
