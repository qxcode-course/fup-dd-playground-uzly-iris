package main
import "fmt"
func main() {
    lista := make([]int, 5)
    
    for i := range lista {
        fmt.Scan(&lista[i])
    }
    x := indenticar_menor(lista)
    fmt.Println(x)
}

func indenticar_menor(lista []int) int {
    var cont, elem int
    for _, elem = range lista {
        for _, i := range lista {
            if i >= elem {
                cont += 1
            } else {
                cont = 0
                break
            }
        }

        if cont == 5 {
            return elem
        }
    }
    return elem
}