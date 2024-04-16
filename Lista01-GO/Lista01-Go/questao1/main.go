package main

import "fmt"

func main() {
    var n1 float64
    var n2 float64
    var n3 float64
    var result float64
  
    fmt.Printf("Digite nota1: ")
    fmt.Scanln(&n1)
    fmt.Printf("Digite nota2: ")
    fmt.Scanln(&n2)
    fmt.Printf("Digite nota3: ")
    fmt.Scanln(&n3)
  
    result = (n1 + n2 + n3) /3
  
    fmt.Printf("MEDIA: %.2f\n", result)

    if result >= 6 {
      fmt.Printf("APROVADO\n")
    } else {
      fmt.Printf("REPROVADO\n")
  }
  
}