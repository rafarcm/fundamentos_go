package enderecos_test

import (
	. "introducao-teste/enderecos" // O . é usado para não precisar ser referenciado (Utilizado mais para testes)
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel() // Roda este teste em paralelo com outros testes que também tem a chamada esta função

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua dos Ingleses", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Estrada das Almas", "Estrada"},
		{"Rodovia Castelo Branco", "Rodovia"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)

		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo de endereço recebido (%s) é diferente do esperado (%s)!", tipoDeEnderecoRecebido, cenario.retornoEsperado)
		}
	}

}
