package main
import "fmt"
func main() {
    var n, cont, div int
    fmt.Scan(&n)
    cont = 0
    div = n
    seg := div

    for i := 2 ; i < n ; {
        
        if div%i == 0 {
            div = div / i
            cont += 1
            if div % i == 0 {
                seg = div
            }
        } else {
            if seg % i == 0 {
                fmt.Println(i, cont)
                seg = div
            }
            cont = 0
            i ++
            if div == 1 {
            break
        }
        }
        
    }
}
