package main
import "fmt"
func main() {
    var n, i int
    fmt.Scan(&n)

    fmt.Print("[ ")
    for i = 0 ; i < 10 ; i ++ {
        if i != n {
            fmt.Print(i, " ")
        }
    }

    if n == 10 {
        fmt.Print("]\n")
    } else {
        fmt.Print("ceu ]\n")
    }
}
