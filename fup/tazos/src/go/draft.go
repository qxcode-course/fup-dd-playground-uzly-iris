package main
import "fmt"
func main() {
    var qtd int
    fmt.Scan(&qtd)
    tazos := make([]int, qtd)
    bckp := make([]int, qtd)
    for i := range tazos {
        fmt.Scan(&tazos[i])
    }

    for _, elem := range tazos {
        

    }
}
