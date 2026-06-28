package main
import "fmt"
func main() {
    var l string
    fmt.Scan(&l)
    c := l[0]

    if c >= 'a' && c <= 'z'{
        fmt.Println(string(c - 32))
    } else if c >= 'A' && c <= 'Z'{
        fmt.Println(string(c + 32))
    } else {
        fmt.Println(string(c))
    }
}