package main

import "fmt"

func main() {
    var x, y int

    fmt.Println("Digite dois números inteiros separados por um espaço:")
    fmt.Scanln(&x, &y)

    if x%2 == 0 {
        for i := 0; i < y; i++ {
            fmt.Print(x, " ")
            x += 2
        }
        fmt.Println() 
    } else {
        fmt.Println("O PRIMEIRO NUMERO NAO E PAR")
    }
}