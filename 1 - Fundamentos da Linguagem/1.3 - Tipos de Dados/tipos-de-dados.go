package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Tipos de Dados")

	// NÚMEROS INTEIROS
	fmt.Println("---------------------------------------------")
	var numero1 int64 = -1000000
	fmt.Println(numero1)

	var numero2 uint8 = 100 // Só recebe inteiros positivos
	fmt.Println(numero2)

	// Alias
	// INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	// UINT8 =  BYTE
	var numero4 byte = 123
	fmt.Println(numero4)
	// FIM NÚMEROS INTEIROS

	// NÚMEROS REAIS
	fmt.Println("---------------------------------------------")
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float32 = 123000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)
	// FIM NÚMEROS REAIS

	// STRINGS
	fmt.Println("---------------------------------------------")
	var str1 string = "Texto 1"
	fmt.Println(str1)

	str2 := "Texto 2"
	fmt.Println(str2)

	// Não existe char no go. Para o exemplo abaixo ele sempre irá imprimir o código ascii do caracter
	char1 := '@'
	fmt.Println(char1)
	// FIM STRINGS

	// BOOLEAN
	fmt.Println("---------------------------------------------")
	var boolean1 bool = true
	fmt.Println(boolean1)

	var boolean2 bool
	fmt.Println(boolean2)

	boolean3 := true
	fmt.Println(boolean3)
	// FIM BOOLEAN

	// ERRORS
	fmt.Println("---------------------------------------------")
	var erro1 error
	fmt.Println(erro1)

	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)
	// FIM ERRORS
}
