package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	fmt.Println("Funções")

	soma := somar(10, 20)
	fmt.Println(soma)

	var f1 = func(txt string) string {
		fmt.Println(txt)
		return "Retorno:" + txt
	}

	resultado := f1("Texto da função 1")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	// Pega o primeiro valor retornado e ignora o segundo
	resultadoSoma1, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma1)

	// ignora o primeiro valor e pega o segundo valor
	_, resultadoSubtracao1 := calculosMatematicos(10, 15)
	fmt.Println(resultadoSubtracao1)
}
