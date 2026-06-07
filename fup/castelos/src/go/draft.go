package main
import "fmt"
func main() {
    var num int
    fmt.Scan(&num)

    quadrado := false

    for i := 1; i*i <= num; i++{
        if i*i == num{
            quadrado = true
            break
        }
    }

    if !quadrado{
        fmt.Println("nao")
    } else{
        fmt.Println("sim")
    }
}