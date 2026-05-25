package main
import "fmt"
func main() {
    var qtd int
    fmt.Scan(&qtd)
    arr := make([]int, qtd)
    for i := range arr {
        fmt.Scan(&arr[i])
    }

    par := make([]int, 0)
    imp := make([]int, 0)

    for _, elem := range arr {
        if elem % 2 == 0 {
            par = append(par, elem)
        } else {
            imp = append(imp, elem)
        }
    }

    fmt.Print("[ ")
    if len(imp) > 0 {
        for _, i := range imp {
            fmt.Print(i, " ")
        }
    }
    fmt.Print("]\n")

    fmt.Print("[ ")
    if len(par) > 0 {
        for _, i := range par {
            fmt.Print(i, " ")
        }
    }
    fmt.Print("]\n")
}
