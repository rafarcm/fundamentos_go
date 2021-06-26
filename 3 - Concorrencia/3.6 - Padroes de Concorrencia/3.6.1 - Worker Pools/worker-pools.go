package main

import "fmt"

func fibonacci(posicao int) int {
	// Condição de parada
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

// O canal tarefas apenas envia dados e o canal resultados apenas recebe dados
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func main() {
	fmt.Println("Worker Pools")

	fmt.Println("---------------------------------------------")

	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	// Chamando várias goroutines, vai deixar o código mais rápido
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 1; i <= 45; i++ {
		tarefas <- i
	}

	close(tarefas)

	for i := 1; i <= 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}
