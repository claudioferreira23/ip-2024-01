package main

import(
	"fmt"
	"sort"
)

func main() {
	x := []int{11, 14, 7, 8, 15, 5}
	ordena(x)
	fmt.Print(x)
}

func ordena(vetor []int) {
	sort.Ints(vetor)
}
