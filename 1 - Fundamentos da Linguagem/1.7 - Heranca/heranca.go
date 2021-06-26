package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa    // Joga todos os campos de pessoa  em estudadnte (tipo Ctrl + c - Ctrl + v)
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança (Só que não)")

	pessoa1 := pessoa{"Mel", "Carolina", 3, 80}
	fmt.Println(pessoa1)

	estudante1 := estudante{pessoa1, "Engenharia", "USP"}
	fmt.Println(estudante1)
	fmt.Println(estudante1.nome)

	estudante2 := estudante{pessoa{"Rafael", "Moras", 36, 180}, "Medicina", "FMU"}
	fmt.Println(estudante2)
}
