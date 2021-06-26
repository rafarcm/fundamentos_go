package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Structs")

	var usuario1 usuario
	fmt.Println(usuario1)

	usuario1.nome = "Rafael"
	usuario1.idade = 36
	fmt.Println(usuario1)

	endereco1 := endereco{"Rua Dos Bobos", 10}
	usuario2 := usuario{"Jaqueline", 30, endereco1}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Mel"}
	fmt.Println(usuario3)

	usuario4 := usuario{idade: 10}
	fmt.Println(usuario4)
}
