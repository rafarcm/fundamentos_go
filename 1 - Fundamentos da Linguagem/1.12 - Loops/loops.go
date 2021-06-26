package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	i := 0
	fmt.Println("---------------------------------------------")
	for i < 5 {
		i++
		fmt.Println("Incrementando i", i)
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	fmt.Println("---------------------------------------------")
	for j := 0; j < 5; j++ {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}

	fmt.Println("---------------------------------------------")
	nomes := [3]string{"Rafael", "Matheus", "Melissa"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	fmt.Println("---------------------------------------------")
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	fmt.Println("---------------------------------------------")
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra) // Imprime os códigos ascii das letras
	}

	fmt.Println("---------------------------------------------")
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra)) // Imprime as letras
	}

	fmt.Println("---------------------------------------------")
	usuario := map[string]string{
		"nome":      "Rafael",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// Loop infinito
	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}
}
