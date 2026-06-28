package main
import "fmt"
func main() {
    var h, mi, d, me, a int
    fmt.Scan(&h, &mi, &d, &me, &a)

    if h < 10 {
        fmt.Printf("0%d", h)
        fmt.Print(":")
    } else {
        fmt.Print(h)
        fmt.Print(":")
    }
    if mi < 10 {
        fmt.Printf("0%d", mi)
    } else {
        fmt.Print(mi)
    }
    fmt.Print(" ")
    if d < 10 {
        fmt.Printf("0%d", d)
        fmt.Print("/")
    } else {
        fmt.Print(d)
        fmt.Print("/")
    }
    if me < 10 {
        fmt.Printf("0%d", me)
        fmt.Print("/")
    } else {
        fmt.Print(me)
        fmt.Print("/")
    }
    if a == 2005 {
        fmt.Println("05")
    } else {
        fmt.Println("88")
    }

}
