package main

import (
	"fmt"
	"math"
)

func distancia(ax, ay, az, bx, by, bz float64) float64 {
	return math.Sqrt(math.Pow(ax-bx, 2) + math.Pow(ay-by, 2) + math.Pow(az-bz, 2))
}

func main() {
	var n int
	fmt.Scan(&n)

	pontos := make([]struct{ x, y, z float64 }, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&pontos[i].x, &pontos[i].y, &pontos[i].z)
	}

	var saida string
	for i := 0; i < n-1; i++ {
		dist := distancia(pontos[i].x, pontos[i].y, pontos[i].z, pontos[i+1].x, pontos[i+1].y, pontos[i+1].z)
		saida += fmt.Sprintf("%.2f\n", dist)
	}

	fmt.Print(saida)
}