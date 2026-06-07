package main
import "fmt"
func main() {
    var valor int
    fmt.Scan(&valor)

    if valor > 0{
        fmt.Println("positivo")
    } else if valor < 0{
        fmt.Println("negativo")
    } else{
        fmt.Println("nulo")
    }
}