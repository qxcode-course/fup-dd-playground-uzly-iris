package main
import "fmt"
func main() {
    var t[3][3] int
    var matriz[3][3] int

    for i := 0 ; i < 3 ; i++ {
        for j := 0 ; j < 3 ; j++ {
            fmt.Scan(&matriz[i][j])
        }
    }

    for j := 0 ; j < 3 ; j++ {
        for i := 0 ; i < 3 ; i++ {
            t[j][i] = matriz[i][j]
        }
    }

    s := true

    for i := 0 ; i < 3 ; i++ {
        for j := 0 ; j < 3 ; j++ {
            if matriz[i][j] != t[i][j] {
                s = false
                break
            }
        }
    }

    if s == true {
        fmt.Println("sim")
    } else {
        fmt.Println("nao")
    }
}