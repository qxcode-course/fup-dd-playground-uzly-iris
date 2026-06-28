package main
import "fmt"
func main() {
    var matriz[5][5] int
    dp := 0
    ds := 0

    for i := 0 ; i < 5 ; i++ {
        for j := 0 ; j < 5 ; j++ {
            fmt.Scan(&matriz[i][j])
        }
    }

    for i := 0 ; i < 5 ; i++ {
        for j := 0 ; j < 5 ; j++ {
            if i == j {
                dp += matriz[i][j]
            }
        }
    }

    for i := 0 ; i < 5 ; i++ {
        ds += matriz[i][5 - 1 - i]
    }

    r := dp - ds
    fmt.Println(r)
}