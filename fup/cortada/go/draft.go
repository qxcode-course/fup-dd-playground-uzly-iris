package main

import "fmt"

func main() {
    var b, t, atpe, atpd int
    fmt.Scan(&b)
    fmt.Scan(&t)

    // lado esquerdo
    atpe = (b + t) * 70 / 2
    
    // lado direito
    atpd = ((160 - b) + (160 - t)) * 70 / 2

    if atpe > atpd {
        fmt.Println(1)
    } else if atpe == atpd {
        fmt.Println(0)
    } else {
        fmt.Println(2)
    }
}
