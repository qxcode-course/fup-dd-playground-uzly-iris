package main
import "fmt"
func main() {
    var x, y int
    fmt.Scan(&x, &y)

    for i := x ; i <= y ; i++ {
        if i%3 == 0 && i%5 == 0{
            fmt.Println("zigzag")
        } else if i%3 == 0{
            fmt.Println("zig")
        } else if i%5 == 0{
            fmt.Println("zag")
        } else {
            fmt.Println(i)
        }
    }
}
