package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Canais")
	fmt.Println("------------------------------")

	canal := make(chan string)

	go escreve("Olá Mundo!", canal)
	fmt.Println("Depois da função escrever começar a ser executada")

	/*for {
		mensagem, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}*/

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")
}

func escreve(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}
