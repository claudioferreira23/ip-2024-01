package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func freqvogal(s string) [5]int {
	var freq [5]int
	for _, caracter := range strings.ToLower(s) {
		switch caracter {
		case 'a':
			freq[0]++
		case 'e':
			freq[1]++
		case 'i':
			freq[2]++
		case 'o':
			freq[3]++
		case 'u':
			freq[4]++
		}
	}
	return freq
}

func diseuclid(v1, v2 [5]int) float64 {
	var soma float64
	for i := 0; i < 5; i++ {
		dif := v1[i] - v2[i]
		soma += float64(dif * dif)
	}
	return math.Sqrt(soma)
}

func main() {
	leitor := bufio.NewReader(os.Stdin)
	entrada, _ := leitor.ReadString('\n')
	entrada = strings.TrimSpace(entrada)

	partes := strings.Split(entrada, ";")
	if len(partes) != 2 {
		fmt.Println("FORMATO INVALIDO!")
		return
	}

	A := strings.TrimSpace(partes[0])
	B := strings.TrimSpace(partes[1])

	freqA := freqvogal(A)
	freqB := freqvogal(B)

	fmt.Printf("(%d,%d,%d,%d,%d)\n", freqA[0], freqA[1], freqA[2], freqA[3], freqA[4])
	fmt.Printf("(%d,%d,%d,%d,%d)\n", freqB[0], freqB[1], freqB[2], freqB[3], freqB[4])

	distance := diseuclid(freqA, freqB)
	fmt.Printf("%.2f\n", distance)
}