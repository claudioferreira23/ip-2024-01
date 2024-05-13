package main

import(
    "fmt"
    "sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	var saidas [][]int

	for i := 0; i < n; i++ {
		var toucas [9]int
		for j := 0; j < 9; j++ {
			fmt.Scan(&toucas[j])
		}

		var saida []int

		for j := 0; j < 9; j++ {
			for k := j + 1; k < 9; k++ {
				soma := 0
				for l := 0; l < 9; l++ {
					if l != j && l != k {
						soma += toucas[l]
					}
				}
				if soma == 100 {
					for l := 0; l < 9; l++ {
						if l != j && l != k {
							saida = append(saida, toucas[l])
						}
					}
					break
				}
			}
			if len(saida) > 0 {
				break
			}
		}
		
		sort.Ints(saida)

		saidas = append(saidas, saida)
	}

	fmt.Println("SAIDA:")
	for _, saida := range saidas {
		for _, n := range saida {
			fmt.Print(n, "\n")
		}
	}
}
