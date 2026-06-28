package main
import "fmt"
func main() {
    var n, mc int
    s := -1

    fmt.Scan(&n)
    matriz := make([][]int, n)

    for i := range matriz {
        matriz[i] = make([]int, n)
        for j := 0 ; j < n ; j++ {
            fmt.Scan(&matriz[i][j])
        }
    }

    for j := 0 ; j < n ; j++ {
        sc := 0
        for i := 0 ; i < n ; i++ {
            e := matriz[i][j]
            sc += e * e
        }

        if j != 0 || sc > s {
            mc = j
        }
    }

    fmt.Println(mc)
}