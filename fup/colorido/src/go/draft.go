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
	fmt.Print("[ ")
    if pe == "e" {
        for i := 0; i < 10; i++ {
		    if i != n && i%2 == 0 {
			    fmt.Print(i, "e ")
		    } else {
                fmt.Print(i, "d ")
            }
	    }
    } else {
        for i := 0; i < 10; i++ {
		    if i != n && i%2 == 0 {
			    fmt.Print(i, "d ")
		    } else {
                fmt.Print(i, "e ")
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