package main

import "fmt"

func main() {
	var h, p, f, d int
	fmt.Scan(&h, &p, &f, &d)

    if d == -1 {
        for i := f;  ; i -- {
            if i == p {
                fmt.Println("N")
                break
            } else if i == h {
                fmt.Println("S")
                break
            } else if i == 0 {
                i += 16
            }
        }
    } else {
        for i := f;  ; i ++ {
            if i == p {
                fmt.Println("N")
                break
            } else if i == h {
                fmt.Println("S")
                break
            } else if i == 15 {
                i -= 16
            }
        }
    }
}
