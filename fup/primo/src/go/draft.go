package main
import (
    "fmt"
    "math"
)
func main() {
    var n int
    fmt.Scan(&n)

    primo := 1

    for i := 2; i <= int(math.Sqrt(float64(n))); i++{
        if n % i == 0{
            primo = 0
            break
        }
    }

    fmt.Println(primo)
}