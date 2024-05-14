package main

import(
	"fmt"
)

func main(){

	var n int

	fmt.Scan(&n)

	x := make([]float64, n, 1000)

	for i := 0; i < n; i++{
		fmt.Scan(&x[i])
	}

	contar := make(map[float64]int)

	for _, num := range x{
		contar[num]++
	}

	for k, _ := range contar{
		fmt.Printf("%.f\n", k)
	}
}