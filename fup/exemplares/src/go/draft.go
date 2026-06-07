package main
import "fmt"
import "sort"

func contains(lista []int, valor int) bool{
    for _, elem := range lista{
        if elem == valor{
            return true
        }
    }
    return false
}

func separar_figurinhas(montante []int) ([]int){

    especies := make([]int, 0)

    for _, animal := range montante {

        if !contains(especies, animal) {
            especies = append(especies, animal)

        }
    }

    return especies

}

func main() {
    var qtd int
    fmt.Scan(&qtd)
    animais := make([]int, qtd)

    for i := range animais{
        fmt.Scan(&animais[i])
    }

    especies := separar_figurinhas(animais)
    sort.Ints(especies)

    for i, elem := range especies{
        if i == len(especies) - 1 {
            fmt.Println(elem)
            break
        } else {
            fmt.Print(elem, " ")
        }
        
    }
}
