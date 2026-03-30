package main

import "fmt"

func main() {
	var n1, n2 int
	var o string
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	fmt.Scan(&o)
	
	if o == "+" {
		soma := n1 + n2
		fmt.Println(soma)
	} else if o == "-" {
		subtração := n1 - n2
		fmt.Println(subtração)
	} else if o == "*" {
		multiplicação := n1 * n2
		fmt.Println(multiplicação)
	} else if o == "/" && n2 != 0 {
		divisão := n1 / n2
		fmt.Println(divisão)
	} else {
		fmt.Println("Não")
	}
}