package main
import "fmt"
func main() {
    var n, cont int 
    fmt.Scan(&n)
    var vet[]int = make([]int, n)

    for i := 0 ; i < n ; i++ {
        fmt.Scan(&vet[i])
    }

    for i := 0 ; i < n - 2 ; i++ {
        if vet[i] == 1 && vet[i + 1] == 0 && vet[i+2] == 0 {
            cont++
        }
    }

    fmt.Println(cont)
}