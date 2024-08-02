package enderecos

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	enderecoEsperado string
}

func TestTipoDeEndereco(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Rua", "Rua"},
		{"rua", "Rua"},
		{"RUA DOS BOBOS", "Rua"},
		{"Praça dos Jardins", "Tipo Inválido"},
		{"", "Tipo Inválido"},
		{" ", "Tipo Inválido"},
		{"Rodovia", "Rodovia"},
		{"rodovia", "Rodovia"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Avenida", "Avenida"},
		{"avenida", "Avenida"},
		{"Avenida Paulista", "Avenida"},
		{"Estrada", "Estrada"},
		{"estrada", "Estrada"},
		{"ESTRADA", "Estrada"},
		{"Estrada dos Montes", "Estrada"},
	}

	for _, value := range cenariosDeTeste {
		tipoDeEndereco := TipoDeEndereco(value.enderecoInserido)
		assert.Equal(t, value.enderecoEsperado, tipoDeEndereco)
	}

	// go test ./... --cover
	// go test --coverprofile cobertura.txt
	// go tool cover --func=cobertura.txt
	// go tool cover --html=cobertura.txt
}
