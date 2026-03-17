package main

import "fmt"

func main() {
    var q1, q2, q3, v1, v2, v3, d, pt, tr float64
    fmt.Scan(&q1)
    fmt.Scan(&q2)
    fmt.Scan(&q3)
    fmt.Scan(&v1)
    fmt.Scan(&v2)
    fmt.Scan(&v3)
    fmt.Scan(&d)
    pt = (q1*v1)+(q2*v2)+(q3*v3)
    tr = d - pt
    fmt.Printf("%.2f\n", tr)
}
