package main

import "fmt"

func main() {
    var njogos int
    fmt.Println("Digite o número de casos de teste:")
    fmt.Scanln(&njogos)

    for i := 1; i <= njogos; i++ {
        var numeroPessoas, percentualPopular, percentualGeral, percentualArquibancada, percentualCadeiras float64
        fmt.Println("Digite o número de pessoas e as percentagens de cada categoria separadas por um espaço:")
        fmt.Scanln(&numeroPessoas, &percentualPopular, &percentualGeral, &percentualArquibancada, &percentualCadeiras)
        
        rendaTotal := (numeroPessoas * ((percentualPopular * 1) + (percentualGeral * 5) + (percentualArquibancada * 10) + (percentualCadeiras * 20))) / 100
        
        fmt.Printf("A RENDA DO JOGO N.%d E = %.2f\n", i, rendaTotal)
    }
}