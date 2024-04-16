package main

import (
    "fmt"
    "math"
)

func main() {
    var altura, aresta float64

    fmt.Println("Digite a altura da pirâmide:")
    fmt.Scanln(&altura)
    fmt.Println("Digite o comprimento da aresta do hexágono:")
    fmt.Scanln(&aresta)

    ab := (3 * math.Pow(aresta, 2) * math.Sqrt(3)) / 2

    volume := (1.0 / 3.0) * ab * altura

    fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS\n", volume)
}
