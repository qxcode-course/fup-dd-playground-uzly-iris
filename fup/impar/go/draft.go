package main

import "fmt"

func main() {
	var p, d1, d2, dt int
	fmt.Scan(&p)
	fmt.Scan(&d1)
	fmt.Scan(&d2)

	dt = d1 + d2

	if p == 0 {
		if dt%2 == 0 {
            fmt.Println(0)
        } else {
            fmt.Println(1)
        }
    } else {
        if dt%2 == 0 {
            fmt.Println(1)
        } else {
            fmt.Println(0)
        }
    }
}
