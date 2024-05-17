package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maiorPremio(numero string, n, d int) string {
	pilha := make([]byte, 0, d)
	remover := n - d

	for i := 0; i < n; i++ {
		for len(pilha) > 0 && remover > 0 && pilha[len(pilha)-1] < numero[i] {
			pilha = pilha[:len(pilha)-1]
			remover--
		}
		pilha = append(pilha, numero[i])
	}

	return string(pilha[:d])
}

func main() {
	leitor := bufio.NewReader(os.Stdin)
	var resultados []string

	for {
		entrada, _ := leitor.ReadString('\n')
		entrada = strings.TrimSpace(entrada)
		if entrada == "0 0" {
			break
		}

		partes := strings.Split(entrada, " ")
		n, _ := strconv.Atoi(partes[0])
		d, _ := strconv.Atoi(partes[1])

		numero, _ := leitor.ReadString('\n')
		numero = strings.TrimSpace(numero)

		resultado := maiorPremio(numero, n, d)
		resultados = append(resultados, resultado)
	}

	for _, resultado := range resultados {
		fmt.Println(resultado)
	}
}