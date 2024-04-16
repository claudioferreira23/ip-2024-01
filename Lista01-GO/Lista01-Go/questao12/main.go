package main

import "fmt"

func main() {
	var tempo int

	fmt.Println("Digite o tempo em horas:")
	fmt.Scanln(&tempo)

	valor := calcularValor(tempo)

	fmt.Printf("O VALOR A PAGAR E = %.2f\n", valor)
}

func calcularValor(tempo int) float64 {
	taxaBase := float64(tempo / 3 * 10)

	horasAdicionais := tempo % 3
	var valorAdicional float64
	if horasAdicionais > 0 {
		valorAdicional = float64(horasAdicionais) * 5
	}

	valorTotal := taxaBase + valorAdicional

	return valorTotal
}