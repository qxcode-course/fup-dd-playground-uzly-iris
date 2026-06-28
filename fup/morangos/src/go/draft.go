package main
import "fmt"
func main() {
    var n1, n2, n3, n4 int
    fmt.Scan(&n1, &n2, &n3, &n4)
    x := n1 * n2
    y := n3 * n4
    if x > y {
        fmt.Println(x)
    } else {
        fmt.Println(y)
    }
}