package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Waitgroup")
	fmt.Println("------------------------------")

	var waitGroup sync.WaitGroup
	waitGroup.Add(3)

	go func() {
		escreve("Goroutine 1!")
		waitGroup.Done() // -1 no contador
	}()

	go func() {
		escreve("Goroutine 2!")
		waitGroup.Done() // -1 no contador
	}()

	go func() {
		escreve("Goroutine 3!")
		waitGroup.Done() // -1 no contador
	}()

	waitGroup.Wait() // Encerra a execução quando o contador chegar em 0
}

func escreve(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
