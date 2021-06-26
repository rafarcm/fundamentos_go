package main

import (
	"fmt"
)

func main() {
	fmt.Println("Canais com Buffer")
	fmt.Println("------------------------------")

	canal := make(chan string, 2) // Defini a capacidade do canal

	canal <- "Canal 1"
	canal <- "Canal 2"

	// Esta linha causaria um deadlock no programa pois ele ficaria esperando um retorno sem ter alguém que o envie
	//canal <- "Canal 3"

	mensagem1 := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem1)
	fmt.Println(mensagem2)
}
