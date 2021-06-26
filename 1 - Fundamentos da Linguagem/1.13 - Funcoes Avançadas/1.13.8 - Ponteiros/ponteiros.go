package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	fmt.Println("Ponteiros")

	fmt.Println("---------------------------------------------")
	numero1 := 20
	numeroInvertido := inverterSinal(numero1)
	fmt.Println(numero1)
	fmt.Println(numeroInvertido)

	numero2 := 40
	fmt.Println(numero2)
	inverterSinalComPonteiro(&numero2)
	fmt.Println(numero2)
}
