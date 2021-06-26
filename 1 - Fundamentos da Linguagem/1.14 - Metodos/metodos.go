package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

// Se você for alterar algum campo do struct você passa a referência do struct
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	fmt.Println("Métodos")

	fmt.Println("---------------------------------------------")
	usuario1 := usuario{"Rafael", 30}
	usuario1.salvar()
	fmt.Println(usuario1.maiorDeIdade())
	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)

	usuario2 := usuario{"Mel", 3}
	usuario2.salvar()
	fmt.Println(usuario2.maiorDeIdade())
	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
