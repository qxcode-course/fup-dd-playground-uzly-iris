package main
import "fmt"
func main() {
    var qtd, B int
    fmt.Scan(&qtd,&B)
    Baruel := make([]int, qtd)
    for i := range Baruel {
        fmt.Scan(&Baruel[i])
    }
    montante := make([]int, qtd)
    sep_fig(montante)
}

func sep_fig(montante []int) { //tupla
    album := make([]int, len(montante))
    repetidas := make([]int, len(montante))
    for _, fig := range montante {
        if contains(album, fig) == false {
            album = append(album, fig)
        } else {
            repetidas = append(repetidas, fig)
        }
    }
    fmt.Print("[ ")
    for _, elem := range repetidas {
        fmt.Print(elem, " ")
    }
    fmt.Println(" ]")

    fmt.Print("[ ")
    for _, elem := range album {
        fmt.Print(elem, " ")
    }
    fmt.Println(" ]")
}

func contains(album []int, fig int) bool {
    for _, elem := range album {
        if elem == fig {
            return true
        }
    }
    return false
}