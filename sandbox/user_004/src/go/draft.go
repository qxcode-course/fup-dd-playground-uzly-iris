package main
import "fmt"
func main() {
    var seios_fartos bool
    fmt.Println("Pode mamar?")
    fmt.Scan(&seios_fartos)

    if seios_fartos == true {
        fmt.Println("gogogo")
    } else {
        fmt.Println("poxa...")
    }
}
