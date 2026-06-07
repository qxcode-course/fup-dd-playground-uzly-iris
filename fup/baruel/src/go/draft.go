package main
import "fmt"

func contains(lista []int, valor int) bool{
    for _, elem := range lista{
        if elem == valor{
            return true
        }
    }
    return false
}

func separar_figurinhas(montante []int) ([]int, []int){

    album := make([]int, 0)
    repetidas := make([]int, 0)

    for _, fig := range montante {

        if !contains(album, fig) {
            album = append(album, fig)

        } else {
            repetidas = append(repetidas, fig)
        }
    }

    return album, repetidas

}

func main() {
    var totalAlbum int
    var totalFigurinhas int

    fmt.Scan(&totalAlbum)
    fmt.Scan(&totalFigurinhas)

    montante := make([]int, totalFigurinhas)

    for i := 0; i < totalFigurinhas; i++{
        fmt.Scan(&montante[i])
    }

    album, repetidas := separar_figurinhas(montante)

    faltando := make([]int, 0)

    for i := 1; i <= totalAlbum; i++ {
        if !contains(album, i){
            faltando = append(faltando, i)
        }
    }

    fmt.Print("[ ")
    for _, num := range repetidas{
        fmt.Print(num, " ")
    }
    fmt.Println("]")

    fmt.Print("[ ")
    for _, num := range faltando{
        fmt.Print(num, " ")
    }
    fmt.Println("]")    
}
