package main

import (
	"fmt"
)

func main() {

		cab := func(tamanho int) map[int]bool {
		x := make(map[int]bool)
		var num int
		for i := 0; i < tamanho; i++ {
			fmt.Scan(&num)

			if _, existe := x[num]; !existe {
				x[num] = true
			} else {
				i--
			}
		}
		return x
	}

	imprimec := func(x map[int]bool) {
		fmt.Print("(")
		primeiro := true
		for key := range x {
			if !primeiro {
				fmt.Print(",")
			}
			fmt.Print(key)
			primeiro = false
		}
		fmt.Println(")")
	}

	var tamanhoA, tamanhoB int

	for {
		fmt.Scan(&tamanhoA)
		if tamanhoA >= 1 && tamanhoA <= 100 {
			break
		}
	}
	for {
		fmt.Scan(&tamanhoB)
		if tamanhoB >= 1 && tamanhoB <= 100 {
			break
		}
	}

	xA := cab(tamanhoA)
	xB := cab(tamanhoB)

	uniao := make(map[int]bool)
	interseccao := make(map[int]bool)

	for k := range xA {
		uniao[k] = true
		if _, existe := xB[k]; existe {
			interseccao[k] = true
		}
	}
	for k := range xB {
		uniao[k] = true
	}

	fmt.Println("União:")
	imprimec(uniao)
	fmt.Println("Interseção:")
	imprimec(interseccao)
}