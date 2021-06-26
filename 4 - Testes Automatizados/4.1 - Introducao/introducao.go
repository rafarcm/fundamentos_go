package main

import (
	"fmt"
	"introducao-teste/enderecos"
)

func main() {
	tipoDeEndereco := enderecos.TipoDeEndereco("Rodovia dos Imigrantes")
	fmt.Println(tipoDeEndereco)
}
