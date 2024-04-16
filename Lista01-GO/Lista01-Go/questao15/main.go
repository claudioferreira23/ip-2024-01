package main

import (
    "fmt"
)

func main() {
    var N int

    fmt.Println("Digite um valor inteiro N (5 < N < 2000):")
    fmt.Scanln(&N)

    if N <= 5 || N >= 2000 {
        fmt.Println("N está fora do intervalo válido.")
        return
    }

    for i := 2; i <= N; i += 2 {
        quadrado := i * i
        fmt.Printf("%dˆ2 = %d\n", i, quadrado)
    }
}
