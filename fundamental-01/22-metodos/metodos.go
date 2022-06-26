package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (user usuario) salvar() {
	fmt.Printf("Salvando user no banco de dados, que tem nome %s \n",
		user.nome)
}

func (user usuario) maiorDeIdade() bool {
	return user.idade >= 18
}

func (user *usuario) fazerAniversario() {
	user.idade++
}

func main() {
	// Principal diferença de métodos para funções, é que ele tem que estar
	// associado a alguma coisa, como struct, interface
	usuario1 := usuario{"Mateus", 23}
	fmt.Println(usuario1)
	usuario1.salvar()
	maiorDeIdade := usuario1.maiorDeIdade()
	fmt.Println(maiorDeIdade)
	usuario1.fazerAniversario()
	fmt.Println(usuario1)
}
