package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1     // Atribuição de valores no go funciona como uma cópia do valor
	fmt.Println(variavel1, variavel2) // Imprime 10 10

	variavel1++
	fmt.Println(variavel1, variavel2) // Só altera o valor da variavel1 (Imprime 11 10)

	// Ponteiro é uma referência de memória
	var variavel3 int = 100
	var ponteiro *int = &variavel3
	fmt.Println(variavel3, ponteiro)  // Imprime 100 e o endereço da memória
	fmt.Println(variavel3, *ponteiro) // Desreferenciação - Imprime 100 100

	variavel3++
	fmt.Println(variavel3, *ponteiro) // Imprime 101 101
}
