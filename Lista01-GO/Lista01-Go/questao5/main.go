package main

import "fmt"

func main() {
    var conta int
    var consumo float64
    var tipoConsumidor rune

    fmt.Println("Digite a conta do cliente, o consumo de água em metros cúbicos e o tipo de consumidor (R - Residencial, C - Comercial, I - Industrial), separados por espaço:")
    fmt.Scan(&conta, &consumo, &tipoConsumidor)

    var valorConta float64

    switch tipoConsumidor {
    case 'R':
        valorConta = 5 + 0.05*consumo
    case 'C':
        if consumo > 80 {
            valorConta = 500 + 0.25*(consumo-80)
        } else {
            valorConta = 500
        }
    case 'I':
        if consumo > 100 {
            valorConta = 800 + 0.04*(consumo-100)
        } else {
            valorConta = 800
        }
    }

    fmt.Printf("CONTA = %d\nVALOR DA CONTA = %.2f\n", conta, valorConta)
}