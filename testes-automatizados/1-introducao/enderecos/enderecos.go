package enderecos

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

// TipoDeEndereco verifica se um endereço tem um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTemTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemTipoValido = true
		}
	}
	if enderecoTemTipoValido {
		return cases.Title(language.Und, cases.NoLower).String(primeiraPalavraDoEndereco)
	}

	return "Tipo Inválido"

}
