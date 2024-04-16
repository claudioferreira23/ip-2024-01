package main

import "fmt"

func main() {
    var n1 int
    var n2 int
    var n3 int

    fmt.Println("Digite o primeiro número:")
    fmt.Scanln(&n1)
    fmt.Println("Digite o segundo número:")
    fmt.Scanln(&n2)
    fmt.Println("Digite o terceiro número:")
    fmt.Scanln(&n3)

    concatenacao := (n1 * 100) + (n2 *10) + n3
    quadrado := concatenacao * concatenacao

    if n1 < 0 || n1 > 9 || n2 < 0 || n2 > 9 || n3 < 0 || n3 > 9 {
        fmt.Printf("DIGITO INVALIDO")
    } else {
        fmt.Printf("%d e %d\n", concatenacao, quadrado)
    }
}
