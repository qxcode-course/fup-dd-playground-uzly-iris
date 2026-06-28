package main
import "fmt"
func main() {
    var nums int
    var matriz[3][3]int

    for i := range matriz{
        for j := range matriz {
            fmt.Scan(&nums)
            matriz[i][j] = nums
        }
    }

    var l[3]int
    var c[3]int
    dp := 0
    ds := 0

    for i := 0 ; i < 3 ; i++ {
        for j := 0 ; j < 3 ; j++ {
            c[j] += matriz[i][j]
        }
    }

    for j := 0 ; j < 3 ; j++ {
        for i := 0 ; i < 3 ; i++ {
            l[i] += matriz[i][j]
        }
    }

    for i := 0 ; i < 3 ; i++ {
        for j := 0 ; j < 3 ; j++ {
            if i == j {
                dp += matriz[i][j]
            }
        }
    }

    for i := 0 ; i < 3 ; i++ {
        ds += matriz[i][3-1-1]
    }

    if l[0] != l[1] && l[1] != l[2] {
        fmt.Println("nao")
    } else if c[0] != c[1] && c[1] != c[2] {
        fmt.Println("nao")
    } else if dp != ds {
        fmt.Println("nao")
    } else {
        fmt.Println("sim")
    }
}