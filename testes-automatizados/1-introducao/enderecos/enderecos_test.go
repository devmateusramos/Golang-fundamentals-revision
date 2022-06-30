// TESTE DE UNIDADE

package enderecos_test

// Pode usar o _test mesmo estando na mesma pasta que tem outro pacote
// Mas vai ter que ficar importando o pacote endereco.TipoDeEndereco
// E pra não ficar pondo algo grande importando do pacote
// você usa o . daí o Go vai enteder que é o pacote principal
// E vai poder usar direto o TipoDeEndereco
import (
	. "endereco/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua abc", "Rua"},
		{"Avenida abc", "Avenida"},
		{"Rodovia palmas", "Rodovia"},
		//{"Campo grande", "Tipo Inválido"},
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

func TestQualquer(t *testing.T) {
	t.Parallel() // rodar em paralelo

	if 1 > 2 {
		t.Errorf("Teste quebrou")
	}
}

// Executar testes de todos os pacotes dentro da pasta que você está no terminal
// use o -> go test ./...
// go test -v é mais descritivo, aparece qual a função de test q foi executada e se passou
// ver cobertura usa o go test --cover
// arquivo txt com relatório de linhas cobertas por teste
// go test --coverprofile nomeDoFile.txt
// o arquivo não é muito legível
// pra isso passa go tool cover --func=nomeDoFile.txt
// esse mostra qual função não está toda coberta
//go tool cover --html=nomeDoFile.txt
// esse é bem descrito
// linhas cinza o que não é testado mesmo
// linha verde o que foi testado
// linha vermelha o que não foi testado
