package main

import "fmt"

func main() {
	var temperatura float64
	var chuva float64

	fmt.Println("Digite a temperatura em Fahrenheit:")
	fmt.Scanln(&temperatura)
	fmt.Println("Digite a quantidade de chuva em polegadas:")
	fmt.Scanln(&chuva)

	celsius := (5 * (temperatura - 32)) / 9
	
	chuvaMilimetros := chuva * 25.4

	fmt.Printf("O VALOR EM CELSIUS = %.2f\n", celsius)
	fmt.Printf("A QUANTIDADE DE CHUVA E = %.2f\n", chuvaMilimetros)
}