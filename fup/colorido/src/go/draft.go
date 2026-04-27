package main

import "fmt"

func main() {
	var n int
    var pe string
	fmt.Scan(&n)
    fmt.Scan(&pe)
	contagem(n, pe)
}


func contagem(n int, pe string) {
	var cont int
	fmt.Print("[ ")
    if pe == "e" {
        for i := 0; i < 10; i++ {
			if i == n {
				continue
			} else {
				if cont%2 == 0 {
					fmt.Print(i, "e ")
					cont += 1
				} else if cont%2 != 0 {
					fmt.Print(i, "d ")
					cont += 1
				}
			}
		}
	} else {
        for i := 0; i < 10; i++ {
			if i == n {
				continue
			} else {
				if cont%2 == 0 {
					fmt.Print(i, "d ")
					cont += 1
				} else if cont%2 != 0 {
					fmt.Print(i, "e ")
					cont += 1
				}
			}
	    }
    }

    pedra_ceu(n)
}	


func pedra_ceu(n int) {
    if n != 10 {
		fmt.Print("ceu ")
	}
	fmt.Println("]")
}