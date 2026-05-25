package main
import "fmt"

func main() {
    var qtd int
    fmt.Scan(&qtd)
    arr := make([]int, qtd)
    for i := range arr {
        fmt.Scan(&arr[i])
    }
    str := 0
    end := len(arr) - 1

    for str < end {
        arr[str], arr[end] = arr[end], arr[str]
        str++
        end--
    }

    fmt.Print("[ ")
    for i := range arr {
        fmt.Print(arr[i], " ")
    }
    fmt.Print("]\n")
}
