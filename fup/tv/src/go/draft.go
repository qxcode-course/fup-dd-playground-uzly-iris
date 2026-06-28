package main
import "fmt"
func main() {
    var v, p float64
    fmt.Scan(&v, &p)
    switch p {
    case 1:
        fmt.Printf("%.2f\n", v)
        fmt.Printf("%.2f\n", v)
    case 2:
        j := (v/100) * 5
        v += j
        j = v/2
        fmt.Printf("%.2f\n", j)
        fmt.Printf("%.2f\n", v)
    case 3:
        j := (v/100) * 10
        v += j
        j = v/3
        fmt.Printf("%.2f\n", j)
        fmt.Printf("%.2f\n", v)
    case 4:
        j := (v/100) * 15
        v += j
        j = v/4
        fmt.Printf("%.2f\n", j)
        fmt.Printf("%.2f\n", v)
    case 5:
        j := (v/100) * 20
        v += j
        j = v/5
        fmt.Printf("%.2f\n", j)
        fmt.Printf("%.2f\n", v)
    case 6:
        j := (v/100) * 25
        v += j
        j = v/6
        fmt.Printf("%.2f\n", j)
        fmt.Printf("%.2f\n", v)
    case 7:
        j := (v/100) * 30
        v += j
        j = v/7
        fmt.Printf("%.2f\n", j)
        fmt.Printf("%.2f\n", v)
    case 8:
        j := (v/100) * 35
        v += j
        j = v/8
        fmt.Printf("%.2f\n", j)
        fmt.Printf("%.2f\n", v)
    case 9:
        j := (v/100) * 40
        v += j
        j = v/9
        fmt.Printf("%.2f\n", j)
        fmt.Printf("%.2f\n", v)
    case 10:
        j := (v/100) * 45
        v += j
        j = v/10
        fmt.Printf("%.2f\n", j)
        fmt.Printf("%.2f\n", v)
    }
}
