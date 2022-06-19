package main

import "fmt"

type pessoa struct {
	idade     uint8
	altura    uint8
	nome      string
	sobrenome string
}

type estudante struct {
	pessoa // a diferença é que aqui ele vai ser acessado .atributosQueTemEmPessoa
	// Ao invés de ser .pessoa.AtributosQueTemEmPessoa
	curso string
}

func main() {
	p1 := pessoa{22, 180, "Mateus", "Ramos"}
	fmt.Println(p1)
	estudante1 := estudante{p1, "Ciência da Computação"}
	fmt.Println(estudante1.nome)
	fmt.Println(estudante1.pessoa.nome)
}
