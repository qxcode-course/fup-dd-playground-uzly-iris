package main

import "fmt"

func main() {
	var c, a int
	fmt.Scan(&c)
	fmt.Scan(&a)

	r := a % (c - 1)

	if r == 0 {
		viagens := a / (c - 1)
		fmt.Println(viagens)
	} else {
		viagens := (a / (c - 1)) + 1
		fmt.Println(viagens)
	}
}
