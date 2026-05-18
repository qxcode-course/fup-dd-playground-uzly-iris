package main

import "fmt"

func main() {
	var q1, q2 int
	v1 := make([]int, q1)
	for i := range v1 {
		fmt.Scan(&v1[i])
	}
	v2 := make([]int, q2)
	for i := range v2 {
		fmt.Scan(&v2[i])
	}
	contains(v1, v2)

}

func contains(v1 []int, v2 []int) {
	cont := 0
	for _, i := range v1{
        for _, elem := range v2{
            if i == elem {
				cont += 1
				break
			}
        }
    }
    if cont > 1 {
        fmt.Println("sim")
    } else {
         fmt.Println("nao")
    }
}