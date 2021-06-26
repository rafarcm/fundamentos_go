package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Goroutines")
	fmt.Println("------------------------------")

	//Não espera a função terminar e já passa para a execução da próxima linha
	go escreve("Olá Mundo!") //goroutine
	escreve("Programando em Go!")

}

func escreve(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
