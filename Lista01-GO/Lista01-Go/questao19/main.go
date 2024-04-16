package main

import (
	"fmt"
)

func calcularSoma(n int) {
	if n <= 1 {
		fmt.Println("Numero invalido!")
		return
	}

	var soma float64 = 0
	for k := 1; k <= n; k++ {
		soma += 1 / float64(k)
	}

	fmt.Printf("%.6f\n", soma)
}

func main() {
	var n int
	fmt.Println("Digite um número inteiro positivo maior que 1:")
	fmt.Scanln(&n)
	calcularSoma(n)
}