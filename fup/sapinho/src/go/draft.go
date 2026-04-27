package main
import "fmt"
func main() {
    var p, s, e,f int
    fmt.Scan(&p, &s, &e)

    for i := 0 ; ; {
        
        if i >= p - s {
            fmt.Println(i, "saiu")
            break
        } else {
            f = i + s
            fmt.Println(i, f)
            i -= e
            i += s
            continue
        }
        
    }
    
}