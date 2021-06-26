package main

import "fmt"

var n int

// É executada antes da função main
// Muito utilizada para setar valores de inicialização antes da execução
func init() {
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Executando a função main")
	fmt.Println(n)
}
