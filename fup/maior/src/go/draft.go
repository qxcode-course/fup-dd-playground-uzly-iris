package main
import "fmt"
func main() {
    var c1, c2 float64
    var m string
    fmt.Scan(&c1, &m, &c2)
    
    if c1 == c2 {
        fmt.Println("primeiro")
    } else if m == "m" && c2 > c1 {
        fmt.Println("primeiro")
    } else if m == "M" && c2 < c1 {
        fmt.Println("primeiro")
    } else {
        fmt.Println("segundo")
    }
}