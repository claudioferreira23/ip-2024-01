package main

import "fmt"

func main() {
	var salarioMinimo float64
	var consumoKW float64

	fmt.Println("Digite o valor do salario minimo:")
	fmt.Scanln(&salarioMinimo)
	fmt.Println("Digite a quantidade de KW que foi consumido:")
	fmt.Scanln(&consumoKW)

	custoPorKW := salarioMinimo * 0.7 / 100

	custoConsumo := consumoKW * custoPorKW

	custoDesconto := custoConsumo * 0.9

	fmt.Printf("Custo por KW: R$ %.2f\n", custoPorKW)
	fmt.Printf("Custo do consumo: R$ %.2f\n", custoConsumo)
	fmt.Printf("Custo com desconto: R$ %.2f\n", custoDesconto)
}