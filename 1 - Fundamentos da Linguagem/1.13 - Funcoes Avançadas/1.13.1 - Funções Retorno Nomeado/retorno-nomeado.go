package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2      // Não preciso colocar a inferência de tipo (:=)
	subtracao = n1 - n2 // Não preciso colocar a inferência de tipo (:=)
	return              // Não preciso declarar as variáveis que serão retornadas
}

func main() {
	fmt.Println("Funções Retorno Nomeado")

	fmt.Println("---------------------------------------------")
	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Println(soma, subtracao)
}
