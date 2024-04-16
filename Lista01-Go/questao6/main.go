package main

import "fmt"

func main() {
    var numero int

    fmt.Println("Digite o número de temperaturas a serem convertidas:")
    fmt.Scanln(&numero)

    for i := 0; i < numero; i++ {
        var fahrenheit float64
		
        fmt.Println("Digite a temperatura em Fahrenheit:")
        fmt.Scanln(&fahrenheit)

        celsius := (5 * (fahrenheit - 32)) / 9

        fmt.Printf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS\n", fahrenheit, celsius)
    }
}