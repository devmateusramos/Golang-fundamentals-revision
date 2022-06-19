package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	// struct é uma coleção de campos
	// cada campo tem um nome e um tipo
	// cria com type, o struct é um tipo
	var u usuario
	fmt.Println(u) // ele vai jogar o valor zero de cada campo
	u.nome = "Mateus"
	u.idade = 22
	fmt.Println(u)
	u2 := usuario{"Mateus", 22, endereco{"Palmas", 17}}
	fmt.Println(u2)
	// Para passar só um valor, deve colocar explícito o campo
	u3 := usuario{nome: "Mateus"}
	fmt.Println(u3)
	endereco1 := endereco{"Palmas", 21}
	u4 := usuario{"Mateus", 22, endereco1}
	fmt.Println(u4)
	fmt.Println("teste herança: ", u2.endereco.numero)
}
