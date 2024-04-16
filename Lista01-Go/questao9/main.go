package main

import "fmt"

func main() {
	var A float64
	var B float64
	var C float64

	fmt.Println("Digite o coeficiente A:")
	fmt.Scanln(&A)
	fmt.Println("Digite o coeficiente B:")
	fmt.Scanln(&B)
	fmt.Println("Digite o coeficiente C:")
	fmt.Scanln(&C)

	Delta := B * B - 4 * A * C

	fmt.Printf("O VALOR DO DELTA E = %.2f\n", Delta)
}