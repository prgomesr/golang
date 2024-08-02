package main

import (
	"fmt"
	"testes-automatizados/enderecos"
)

func main() {
	tipoDeEndereco := enderecos.TipoDeEndereco("rua da primavera")
	fmt.Println(tipoDeEndereco)
}
