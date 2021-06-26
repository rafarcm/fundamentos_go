package main

import "fmt"

func main() {
	fmt.Println("Operadores")

	// ARITMÉTICOS
	fmt.Println("---------------------------------------------")
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 13 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)
	// FIM ARITMÉTICOS

	// ATRIBUIÇÃO
	fmt.Println("---------------------------------------------")
	var variavel1 string = "String 1"
	variavel2 := "String 2"

	fmt.Println(variavel1, variavel2)
	// FIM ATRIBUIÇÃO

	// OPERADORES RELACIONAIS
	fmt.Println("---------------------------------------------")
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)
	// FIM OPERADORES RELACIONAIS

	// OPERADORES LÓGICOS
	fmt.Println("---------------------------------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && verdadeiro)
	fmt.Println(verdadeiro && falso)
	fmt.Println(falso && falso)
	fmt.Println(verdadeiro || verdadeiro)
	fmt.Println(verdadeiro || falso)
	fmt.Println(falso || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	// FIM OPERADORES LÓGICOS

	// OPERADORES UNÁRIOS
	fmt.Println("---------------------------------------------")
	numero := 10
	fmt.Println(numero)
	numero++
	fmt.Println(numero)
	numero += 15 // numero = numero + 15
	fmt.Println(numero)
	numero--
	fmt.Println(numero)
	numero -= 3 // numero = numero -3
	fmt.Println(numero)
	numero *= 4 // numero = numero * 4
	fmt.Println(numero)
	numero /= 10 // numero = numero / 10
	fmt.Println(numero)
	numero %= 9 // numero = numero % 9
	fmt.Println(numero)
	// FIM OPERADORES UNÁRIOS
}
