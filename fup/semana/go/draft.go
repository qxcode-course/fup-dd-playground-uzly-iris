package main

import "fmt"

func main() {
	var dia, hora int
	fmt.Scan(&dia)
	fmt.Scan(&hora)

	if dia == 1 {
		fmt.Println("NAO")
	} else if dia < 7 {
		if 8 <= hora && hora <= 11 || 14 <= hora && hora <= 17 {
			fmt.Println("SIM")
		} else {
			fmt.Println("NAO")
		}
	} else {
		if 8 <= hora && hora <= 11 {
			fmt.Println("SIM")
		} else {
			fmt.Println("NAO")
		}
	}
}
