package main

import "fmt"

func main() {
	fmt.Println("Funções Anônimas")

	fmt.Println("---------------------------------------------")

	func() {
		fmt.Println("Olá Mundo")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Olá Mundo Parâmetro")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Olá Mundo Parâmetro")

	fmt.Println(retorno)
}
