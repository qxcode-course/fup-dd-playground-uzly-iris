package main

import "fmt"

func main() {
    var tc, tf float64
    fmt.Scan(&tc)
    tf = tc*1.8+32
    fmt.Printf("%.6f\n", tf)
}
