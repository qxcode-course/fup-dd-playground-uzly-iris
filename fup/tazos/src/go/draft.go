package main
import "fmt"
func main() {
    var qtd, bckp int
    fmt.Scan(&qtd)
    tazos := make([]int, qtd)
    for i := range tazos {
        fmt.Scan(&tazos[i])
    }

    for _, elem := range tazos {
        if elem == bckp {
            
        }

    }
}
