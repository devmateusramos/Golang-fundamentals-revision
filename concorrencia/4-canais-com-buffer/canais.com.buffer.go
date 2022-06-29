package main

import "fmt"

func main() {
	canal := make(chan string, 3)
	canal <- "Olá mundo"
	canal <- "Olá mundo"
	// Como estou esperando ele receber tudo na mesma função
	// vai dar deadlock, pois ele vai ficar esperando sem parar recebimentos
	// Enquanto a função é executada
	// Pra isso passamos o buffer, que vai definir quantos recebimentos tem
	// Mas se criarmos algo além da quantidade de buffer, vai travar novamente
	// em deadlock
	// O buffer define até quantas podem entrar, não utilizando o recebimento do
	// canal acima dele, ele não vai ficar esperando, fica sob demanda
	// sem gerar o deadlock
	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
