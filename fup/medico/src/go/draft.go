package main
import "fmt"
func main() {
    var qnt int
    fmt.Scan(&qnt)

    sequencia := make([]int, qnt)
    for i := 0; i < qnt; i++{
        fmt.Scan(&sequencia[i])
    }

    contador := 0


    for i := 0; i < len(sequencia); i++{
        if sequencia[i] == 0{
            temUmPerto := false

            if i > 0 && sequencia[i-1] == 1{
                temUmPerto = true
            }

            if i < len(sequencia)-1 && sequencia[i+1] == 1{
                temUmPerto = true
            }
            if temUmPerto==false{
            contador++
            }
        }

    }

    fmt.Println(contador)
}