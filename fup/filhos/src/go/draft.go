package main

import "fmt"

func main() {
	var idade, quantidade int
	fmt.Scan(&idade)
	fmt.Scan(&quantidade)
    x := 1

	for i := idade; x <= quantidade; i += 2 {
		x += 1
		fmt.Println(i)
	}
}
