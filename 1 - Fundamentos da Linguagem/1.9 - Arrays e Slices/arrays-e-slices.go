package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	// ARRAY
	// Array obrigatoriamente precisa ser identificado o tamanho do array (Posições fixas)
	fmt.Println("---------------------------------------------")
	var array1 [5]int
	fmt.Println(array1)

	array1[0] = 10
	fmt.Println(array1)

	array2 := [6]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5", "Posição 6"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5} // Define o tamanho do array com a quantidade de itens na inicialização
	fmt.Println(array3)
	// FIM ARRAY

	// SLICE
	fmt.Println("---------------------------------------------")
	slice1 := []int{10, 11, 12, 13}
	fmt.Println(slice1)
	slice1 = append(slice1, 14)
	fmt.Println(slice1)

	slice2 := array2[1:4] // Pega o pedaço do Array da posição 1 (inclusivo) até 4 (exclusivo, pega até a posição imediatamente anterior)
	fmt.Println(slice2)

	array2[1] = "Posição 2 Alterada"
	fmt.Println(slice2) // Imprime a posição 2 alterada. Slice funciona como um ponteiro
	// FIM SLICE

	fmt.Println(reflect.TypeOf(array1), reflect.TypeOf(slice1)) // Slice não é a mesma coisa que Array. São tipos diferentes

	// ARRAY INTERNO
	fmt.Println("---------------------------------------------")
	slice3 := make([]float32, 10, 15) // Cria um array interno de 15 posições e retorna um slice com as 10 primeiras posições
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Imprime o tamanho do slice
	fmt.Println(cap(slice3)) // Imprime a capacidade do slice

	fmt.Println("---------------------------------------------")
	slice4 := make([]float32, 5, 6)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	slice4 = append(slice4, 10)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	slice4 = append(slice4, 20) // Para não estourar o slice, o go cria um novo array interno com o dobro da capacidade
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // Imprime 7
	fmt.Println(cap(slice4)) // Imprime 12

	fmt.Println("---------------------------------------------")
	slice5 := make([]float32, 5) // Neste caso a capacidade máxima será igual ao tamanho do array interno
	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))
	// FIM ARRAY INTERNO
}
