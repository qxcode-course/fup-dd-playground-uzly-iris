package main
import "fmt"
func main() {
    var numA, numB int
    fmt.Scan(&numA, &numB)

    if numA>numB{
        fmt.Println(numA)
    } else{
        fmt.Println(numB)
    }
}