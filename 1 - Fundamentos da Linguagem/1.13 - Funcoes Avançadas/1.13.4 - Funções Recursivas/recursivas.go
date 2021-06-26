package main

import "fmt"

func fibonacci(posicao uint) uint {
	// Condição de parada
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções Recursivas")

	fmt.Println("---------------------------------------------")

	// 1 1 2 3 5 8 13 .....

	posicao := uint(7)
	fmt.Println(fibonacci(posicao))

	for i := uint(1); uint(i) <= posicao; i++ {
		fmt.Print(fibonacci(i), " ")
	}
}
