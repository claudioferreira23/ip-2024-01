package main

import "fmt"

func inverte(arr []int, comeco, fim int) {
	if comeco >= fim {
		return
	}

	arr[comeco], arr[fim] = arr[fim], arr[comeco]

	inverte(arr, comeco+1, fim-1)
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

	fmt.Println("Array original:", array)

	inverte(array, 0, len(array)-1)

	fmt.Println("Array invertido:", array)
}
