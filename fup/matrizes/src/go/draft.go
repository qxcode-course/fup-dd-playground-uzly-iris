package main
import "fmt"
func main() {
    var l, c int
    fmt.Scan(&l, &c)
    matriz := make([][]int, l)
    matriz1 := make([][]int, l)
    s := make([][]int, l)

    for i := 0 ; i < l ; i++ {
        matriz[i] = make([]int, c)
        matriz1[i] = make([]int, c)
        s[i] = make([]int, c)
    }

    for i := 0 ; i < l ; i ++ {
        for j := 0 ; j < c ; j++ {
            fmt.Scan(&matriz[i][j])
        }
    }

    for i := 0 ; i < l ; i ++ {
        for j := 0 ; j < c ; j++ {
            fmt.Scan(&matriz1[i][j])
        }
    }

    for i := 0 ; i < l ; i++ {
        fmt.Print("[ ")
        for j := 0 ; j < c ; j++ {
            s[i][j] = matriz[i][j] + matriz1[i][j]
            fmt.Print(s[i][j], " ")
        }
        fmt.Print("]\n")
    }
}