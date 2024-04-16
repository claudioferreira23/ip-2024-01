package main

import "fmt"


func main() {
	var raio, altura float64

	fmt.Println("Digite o raio da lata em metros:")
	fmt.Scanln(&raio)

	fmt.Println("Digite a altura da lata em metros:")
	fmt.Scanln(&altura)

	const pi = 3.14159

	areaCirculo := pi * raio * raio

	areaLateral := 2 * pi * raio * altura

	areaTotal := 2 * areaCirculo + areaLateral

	custo := areaTotal * 100

	fmt.Printf("O VALOR DO CUSTO E = %.2f\n", custo)
}