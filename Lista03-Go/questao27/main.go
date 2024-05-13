package main

import (
	"fmt"
)

func main() {

	intercala := func(q1, q2 int, v1, v2 []int) []int {
		var resultado []int
		i, j := 0, 0
		
		for i < q1 && j < q2 {
			if v1[i] <= v2[j] {
				resultado = append(resultado, v1[i])
				i++
			} else {
				resultado = append(resultado, v2[j])
				j++
			}
		}

		for i < q1 {
			resultado = append(resultado, v1[i])
			i++
		}

		for j < q2 {
			resultado = append(resultado, v2[j])
			j++
		}

		return resultado
	}

	var q1, q2 int
	fmt.Scan(&q1)
	fmt.Scan(&q2)

	v1 := make([]int, q1)
	v2 := make([]int, q2)

	for i := 0; i < q1; i++ {
		fmt.Scan(&v1[i])
	}

	for i := 0; i < q2; i++ {
		fmt.Scan(&v2[i])
	}

	result := intercala(q1, q2, v1, v2)
	fmt.Println("Valores intercalados e ordenados:")
	for _, num := range result {
		fmt.Print(num, "\n")
	}
}