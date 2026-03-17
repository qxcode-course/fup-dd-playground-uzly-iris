package main

import "fmt"

func main() {
    var tempo, h, m, s int
    fmt.Scan(&tempo)
    h = tempo/3600
    m = (tempo%3600)/60
    s = (tempo%3600)%60
    fmt.Printf("%v:%v:%v\n", h, m, s)
}
