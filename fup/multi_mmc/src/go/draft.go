package main

import "fmt"

// Função para calcular o Máximo Divisor Comum (MDC) usando o Algoritmo de Euclides
func mdc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Função para calcular o Mínimo Múltiplo Comum (MMC) entre dois números
func mmcDoisNumeros(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	// Fórmula: MMC(a, b) = (a * b) / MDC(a, b)
	return (a * b) / mdc(a, b)
}

// Função para calcular o MMC de uma lista (slice) de inteiros
func mmcDaLista(numeros []int) int {
	if len(numeros) == 0 {
		return 0
	}

	resultado := numeros[0]
	for i := 1; i < len(numeros); i++ {
		resultado = mmcDoisNumeros(resultado, numeros[i])
	}
	return resultado
}

func main() {
	// Exemplo de lista
	lista := []int{2, 3, 5, 7, 10}
	
	resultadoFinal := mmcDaLista(lista)
	
	fmt.Printf("O MMC da lista %v é: %d\n", lista, resultadoFinal)
}