package main
import "fmt"
func main() {
    var id, di int
    fmt.Scan(&id)

    for i := id; i > 0 ; {
        di *= 10
        di += i % 10
        i = i/10 
    }

    if id == di {
        fmt.Println(1)
    } else {
        fmt.Println(0)
    }
}