package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)
    x := float64(a)
    y := float64(b)
    z := x / y

    fmt.Println(a + b)
    fmt.Println(a - b)
    fmt.Println(a * b)
    fmt.Printf("%.2f\n", z)
    fmt.Println(a % b)
}
