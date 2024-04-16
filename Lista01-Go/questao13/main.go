package main

import "fmt"

func main() {
	var nota float64

	fmt.Println("Digite a nota:")
	fmt.Scanln(&nota)

	if nota >= 9 && nota <= 10 {
		fmt.Printf("NOTA = %.1f CONCEITO = A\n", nota) 
	} else if nota >= 7.5 && nota < 9 {
		fmt.Printf("NOTA = %.1f CONCEITO = B\n", nota)
	} else if nota >= 6 && nota < 7.5 {
		fmt.Printf("NOTA = %.1f CONCEITO = C\n", nota)
	} else if nota >= 0 && nota < 6 {
		fmt.Printf("NOTA = %.1f CONCEITO = D\n", nota)
	}
}

