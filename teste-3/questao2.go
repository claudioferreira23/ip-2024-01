package main

import (
	"bytes"
	"fmt"
)

func troca(vetor []int, i, j int) {
	vetor[i], vetor[j] = vetor[j], vetor[i]
}

func trocaOpostosSeMenor(vetor []int, n int) {
	for i := 0; i < n/2; i++ {
		if vetor[i] < vetor[n-1-i] {
			troca(vetor, i, n-1-i)
		}
	}
}

func main() {
	var casos int
	fmt.Scan(&casos)

	var buffer bytes.Buffer

	for c := 0; c < casos; c++ {
		var n int
		fmt.Scan(&n)

		vetor := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&vetor[i])
		}

		trocaOpostosSeMenor(vetor, n)

		for i := 0; i < n; i++ {
			if i > 0 {
				buffer.WriteString(" ")
			}
			buffer.WriteString(fmt.Sprintf("%d", vetor[i]))
		}
		buffer.WriteString("\n")
	}

	fmt.Print(buffer.String())
}
