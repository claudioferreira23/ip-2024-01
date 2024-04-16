package main

import "fmt"

func main() {
	var A float64
	var B float64
	var C float64
	var D float64

	fmt.Println("Digite o elemento A da matriz:")
	fmt.Scanln(&A)
	fmt.Println("Digite o elemento B da matriz:")
	fmt.Scanln(&B)
	fmt.Println("Digite o elemento C da matriz:")
	fmt.Scanln(&C)
	fmt.Println("Digite o elemento D da matriz:")
	fmt.Scanln(&D)

	Determinante := A * D - B * C

	fmt.Printf("O VALOR DO DETERMINANTE E = %.2f\n", Determinante)
}