package main

import (
	"fmt"
	"math"
)

func main() {
	var T int
	fmt.Scan(&T)

	results := make([]string, T)

	for t := 0; t < T; t++ {
		var N int
		fmt.Scan(&N)

		sequence := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&sequence[i])
		}

		minDistance := math.MaxInt32
		comparisons := 0

		// Calcula a distância entre os elementos adjacentes na sequência
		for i := 1; i < N; i++ {
			for j := 0; j < i; j++ {
				dist := abs(sequence[i] - sequence[j])
				comparisons++
				if dist < minDistance {
					minDistance = dist
				}
			}
		}

		results[t] = fmt.Sprintf("%d %d", minDistance, comparisons)
	}

	// Imprime todos os resultados no final
	for _, result := range results {
		fmt.Println(result)
	}
}

// Função auxiliar para calcular o valor absoluto de um número inteiro
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}