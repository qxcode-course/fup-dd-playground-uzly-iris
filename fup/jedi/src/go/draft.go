package main
import "fmt"
func main() {
    var qtd, cj, cs int
    fmt.Scan(&qtd)
    arr := make([]int, qtd)
    for i := range arr {
        fmt.Scan(&arr[i])
    }

    for i, e := range arr {
        if i < qtd / 2 {
            cj += e
        } else {
            cs += e
        }
    }
    
    if cj > cs {
        fmt.Println("Jedi")
    } else if cj == cs {
        fmt.Println("Empate")
    } else {
        fmt.Println("Sith")
    }
}