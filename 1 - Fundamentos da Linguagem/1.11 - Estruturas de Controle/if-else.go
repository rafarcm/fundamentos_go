package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle - If")

	fmt.Println("---------------------------------------------")
	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual 15")
	}

	// IF INIT
	if outroNumero := numero; outroNumero > 0 { // A variável (outroNumero) só fica disponível no escopo do if. Fora do if ela não existe
		fmt.Println("Maior que 0")
	} else {
		fmt.Println("Menor que 0")
	}
}
