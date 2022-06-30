// TESTE DE UNIDADE

package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua abc", "Rua"},
		{"Avenida abc", "Avenida"},
		{"Rodovia palmas", "Rodovia"},
		{"Campo grande", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("Era esperado %s, mas foi retornado %s",
				tipoDeEnderecoRecebido,
				cenario.retornoEsperado)
		}
	}
}
