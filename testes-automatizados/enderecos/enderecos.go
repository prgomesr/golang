package enderecos

import "strings"

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"Rua", "Avenida", "Estrada", "Rodovia"}

	primeiraPalavraDoEndereco := strings.Split(endereco, " ")[0]

	for _, value := range tiposValidos {
		if strings.ToLower(primeiraPalavraDoEndereco) == strings.ToLower(value) {
			return value
		}
	}

	return "Tipo Inválido"
}
