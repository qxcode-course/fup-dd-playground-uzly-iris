package main
import "fmt"
func main() {
    var l, c int
    fmt.Scan(&l, &c)
    matriz := make([][]int, l)

    for i := 0; i < l ; i ++ {
        matriz[i] = make([]int, c)
    }
    
    for i := 0 ; i < l ; i++ {
        for j := 0 ; j < c ; j++ {
            fmt.Scan(&matriz[i][j])
        }
    }

    sn := 0

    for j := 0 ; j < c ; j++ {
        for i := 0 ; i < l - 1 ; i++ {
            if matriz[i + 1][j] < matriz[i][j]{
                sn++
            }
        }
    }

    fmt.Println(sn)
}