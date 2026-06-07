package main
import "fmt"
func main() {
    var x float64
    fmt.Scan(&x)

    cedulas := []float64 {100, 50, 20, 10, 5, 2}
    moedas := []float64 {1, 0.50, 0.25, 0.10, 0.5}
    quantia := make([]float64, 0)

    for i := 0; ; {
        if x > 1 {
            x = x/cedulas[i]
        } 
    }
}
