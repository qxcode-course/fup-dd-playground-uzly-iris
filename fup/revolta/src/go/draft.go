package main
import "fmt"
func main() {
    var t int
    fmt.Scan(&t)

    num := make([]int, t)
    for i := 0; i < t; i++{
        fmt.Scan(&num[i])
    }

    r := 0
    s := 0

    for i := 0; i < t; i++{
        if num[i] % 2 == 0{
            r += num[i]
        } else {
            s += num[i]
        }
    }
    if r > s{
        fmt.Println("rebeldes")
    } else if s > r{
        fmt.Println("soldados")
    } else {
        fmt.Println("empate")
    }
}
