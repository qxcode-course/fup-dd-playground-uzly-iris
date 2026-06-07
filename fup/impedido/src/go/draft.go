package main
import "fmt"
func main() {
    var jogL, jogR, jogD int
    fmt.Scan(&jogL, &jogR, &jogD)

    if jogR > 50 && jogL < jogR && jogR > jogD{
        fmt.Println("S")
    } else{
        fmt.Println("N")
    }
}