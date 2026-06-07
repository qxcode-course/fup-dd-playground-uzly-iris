package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    v := make([]int, n)

    for i := 0; i < n; i++ {
        fmt.Scan(&v[i])
    }

    //guarda a maior quantidade de repetições encontrada
    //exemplo: se um número apareceu 3 vezes:
    //maior = 3
    maior := 0

    //cria um vetor vazio
    //vai guardar quais números estiveram a maior repetição
    repetidos := []int{}

    //conta quantas vezes o número atual apareceu
    //começa no 1, porque um número sozinho já apareceu uma vez
    cont := 1

    for i := 1; i < n; i++ {

        //verifica se o número atual é igual ao anterior
        //exemplo: 1 1
        //sim?
        //aumenta contagem em 1
        if v[i] == v[i-1] {
            cont++

        //acontece quando o número que está sendo avaliado mudar    
        } else {

            //verifica se encontramos um grupo maior que todos os outros
            if cont > maior {

                //atualiza o recorde de repetições
                maior = cont

                //cria uma nova lista com os vencedores
                //usa v[i-1] porque o grupo que terminou foi o anterior
                repetidos = []int{v[i-1]}

            //verifica se houve empate
            } else if cont == maior {

                //adiciona esse número na lista
                repetidos = append(repetidos, v[i-1])
            }

            //reinicia a contagem, porque um novo número começou
            cont = 1
        }
    }

    //trata o último grupo
    //como o else só acontece quando o número muda, então ele não é análisado
    if cont > maior {
        //se o ultimo grupo foi o maior: ele vira o novo vencedor
        repetidos = []int{v[n-1]}

    //se tiver empatado com o maior
    } else if cont == maior {
        //adiciona o último número a lista
        repetidos = append(repetidos, v[n-1])
    }

    fmt.Print("[ ")
    for _, x := range repetidos {
        fmt.Print(x, " ")
    }
    fmt.Println("]")
}