package main
import "fmt"
func main() {
    var n1, n2 int
    fmt.Scan(&n1, &n2)
    fmt.Println(n1/n2)
    fmt.Println(n1%n2)
    x := float64(n1)
    y := float64(n2)
    z := (x / y)
    fmt.Printf("%.2f\n", z)
}
