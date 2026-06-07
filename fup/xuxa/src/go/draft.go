package main 
import "fmt"
func main() { 
    str := make([]string, 0)

    for i := range str {
        fmt.Scan(&str[i])
    }

    fmt.Println(str)
}