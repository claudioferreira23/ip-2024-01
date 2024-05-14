package main

import (
	"fmt"
	"sort"
)

func main() {
	var saida string
	var n int
	for {
		fmt.Scanln(&n)
		if n == 0 {
			break
		}
		var x []int
		for i := 0; i < n; i++ {
			var num int
			fmt.Scan(&num)
			x = append(x, num)
		}
		M := x[0]
		for _, num := range x {
			if num > M {
				M = num
			}
		}
		sort.Ints(x)
		freq := make([]int, M+1)
		for _, num := range x {
			freq[num]++
		}
		contar := 0
		for i := 0; i <= M; i++ {
			contar += freq[i]
			saida += fmt.Sprintf("(%d) %d\n", i, contar)
		}
	}
	fmt.Print(saida)
}