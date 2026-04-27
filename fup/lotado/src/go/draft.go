package main
import "fmt"
func main() {
    var c, m int
    fmt.Scan(&c)
    for i := 0 ; ; {
        fmt.Scan(&m)
        i += m
        if i == 0 {
            fmt.Println("vazio")
            continue
        } else if i < c {
            fmt.Println("ainda cabe")
            continue
        } else if i == c || i < c*2 {
            fmt.Println("lotado")
            continue
        } else if i >= c*2 {
            fmt.Println("hora de partir")
            break
        }
    }

}
