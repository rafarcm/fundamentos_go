package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÈDIA É EXATAMENTE 6!") // Mata a execução do programa tudo que devria ser executado depois do panic não será executado
}

func main() {
	fmt.Println("Panic e Recover")

	fmt.Println("---------------------------------------------")
	fmt.Println(alunoEstaAprovado(7, 6))
	fmt.Println("Pós execução!")

	fmt.Println("---------------------------------------------")
	fmt.Println(alunoEstaAprovado(3, 4))
	fmt.Println("Pós execução!")

	fmt.Println("---------------------------------------------")
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós execução!")
}
