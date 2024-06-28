package main

import (
	"fmt"
)

func main() {
	var frases []string
	var entrada string

	fmt.Println("Digite as frases (Digite # para parar):")
	for {
		fmt.Scanln(&entrada)
		if entrada == "#" {
			break
		}
		frases = append(frases, entrada)
	}

	fmt.Println("\nFrases invertidas:")
	for _, frase := range frases {
		fmt.Println(inverterString(frase))
	}
}

func inverterString(s string) string {
	caracteres := []rune(s)
	for i, j := 0, len(caracteres)-1; i < j; i, j = i+1, j-1 {
		caracteres[i], caracteres[j] = caracteres[j], caracteres[i]
	}
	return string(caracteres)
}
