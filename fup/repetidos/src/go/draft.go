package main
import "fmt"
func main() {
    var p, n int
    fmt.Scan(&p, &n)
    var lista []int = make([]int, n)
    
    for i := range lista {
        fmt.Scan(&lista[i])
    }
    cont := count(lista, p)
    fmt.Println(cont)
}

func count(lista []int, p int) int {
    contador := 0
    for _, elem := range lista {
        if elem == p {
            contador += 1
        }
    }
    return contador
}