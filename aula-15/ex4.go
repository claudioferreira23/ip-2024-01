package main

import (
	"fmt"
)

func decimalparabinario(n int) string {
	if n == 0 {
		return "0"
	}
	
	binario := decimalparabinario(n / 2)

	binario = binario + fmt.Sprintf("%d", n%2)
	return binario
}

func arrayparabinario(arr []int, resultado []string, indice int) {
	if indice >= len(arr) {
		return
	}

	repbinaria := decimalparabinario(arr[indice])
	resultado[indice] = repbinaria

	arrayparabinario(arr, resultado, indice+1)
}

func main() {
	var array []int
	var n int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		array = append(array, num)
	}

	arraybinario := make([]string, len(array))

	arrayparabinario(array, arraybinario, 0)

	fmt.Println("Array original:", array)
	fmt.Println("Array em binário:", arraybinario)
}