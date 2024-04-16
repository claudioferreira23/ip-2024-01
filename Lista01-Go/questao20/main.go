package main

import "fmt"

func main() {
	var tempoHora, tempoMinuto, tempoSegundo int

	fmt.Println("Digite o valor do tempo em horas")
	fmt.Scanln(&tempoHora)
	fmt.Println("Digite o valor do tempo em minutos")
	fmt.Scanln(&tempoMinuto)
	fmt.Println("Digite o valor do tempo em segundos")
	fmt.Scanln(&tempoSegundo)

	var conversaoSegundos int

	horaSegundos := tempoHora * 3600
	minutoSegundos := tempoMinuto * 60

	conversaoSegundos = horaSegundos + minutoSegundos + tempoSegundo

	fmt.Printf("O TEMPO EM SEGUNDOS E = %d\n", conversaoSegundos)
}