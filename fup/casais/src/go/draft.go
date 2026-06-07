package main
import "fmt"
func main() {
	var qnt int
	fmt.Scan(&qnt)

	contagem := make(map[int]int)
	for i := 0; i < qnt; i++ {
		var animal int
		fmt.Scan(&animal)
		contagem[animal]++
	}

	casais := 0

	for especie, qtdMachos := range contagem {
		if especie > 0 {
			qtdFemeas := contagem[-especie]
			if qtdMachos < qtdFemeas {
				casais += qtdMachos
			} else {
				casais += qtdFemeas
			}
		}
	}
	fmt.Println(casais)
}