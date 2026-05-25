package main
import "fmt"
import "slice"

func main() {
    var max, qtdb int
    fmt.Scan(&max, &qtdb)
    montante := make([]int, qtdb)

    for i := range montante {
        fmt.Scan(montante[i])
    }  

    sepfig(montante)
}

func sepfig(montante []int) {
    album := make([]int, 0)
    repet := make([]int, 0)

    for _, fig := range montante {
        if !Slices(album, fig) {
            album = append(album, fig)
        } else {
            repet = append(repet, fig)
        }
    }

    
}