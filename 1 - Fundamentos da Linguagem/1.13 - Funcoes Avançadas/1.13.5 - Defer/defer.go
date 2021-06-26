package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado!") // Vai executar essa linha imediatamente antes do return
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

// Defer é muito útil para não replicar código
func main() {
	fmt.Println("Defer")

	fmt.Println("---------------------------------------------")
	funcao1()
	funcao2()

	fmt.Println("---------------------------------------------")
	defer funcao1() // DEFER = ADIAR - Será a última linha a ser executada
	funcao2()

	fmt.Println(alunoEstaAprovado(7, 8))
}
