package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) { // Só pode passar um parâmetro variático por função e ele deve ser sempre o último
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println("Funções Variáticas")

	fmt.Println("---------------------------------------------")
	totalDaSoma := soma(10, 10, 20)
	fmt.Println(totalDaSoma)

	totalDaSoma2 := soma(10, 10, 20, 30)
	fmt.Println(totalDaSoma2)

	totalDaSoma3 := soma()
	fmt.Println(totalDaSoma3)

	escrever("Olá Mundo: ", 1, 2, 3, 4, 5, 6, 7, 8)
}
