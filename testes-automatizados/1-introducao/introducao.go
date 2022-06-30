package main

import (
	"endereco/enderecos"
	"fmt"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Avenida arne palmas")
	fmt.Println(tipoEndereco)
}
