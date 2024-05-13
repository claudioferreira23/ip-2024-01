package main

import (
	"fmt"
	"strings"
)

func countingSort(arr []int) {
	n := len(arr)

	max := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	contar := make([]int, max+1)

	for i := 0; i < n; i++ {
		contar[arr[i]]++
	}

	for i := 1; i <= max; i++ {
		contar[i] += contar[i-1]
	}

	ord := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		ord[contar[arr[i]]-1] = arr[i]
		contar[arr[i]]--
	}

	copy(arr, ord)
}

func main() {
	var saidas []string

	for {
		var n int
		fmt.Scan(&n)
		if n == 0 {
			break
		}

		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&arr[i])
		}

		countingSort(arr)

		saida := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arr)), " "), "[]")
		saidas = append(saidas, saida)
	}

	fmt.Println(strings.Join(saidas, "\n"))
}