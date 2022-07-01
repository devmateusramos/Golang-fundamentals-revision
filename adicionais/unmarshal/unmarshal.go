package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome string `json:"nome"` // Se não quiser q converta deixa um - que
	// vai ficar o valor 0 no lugar
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroEmJson := `{"nome":"Lucian","idade":3,"raca":"Pastor Alemão"}`
	fmt.Println(cachorroEmJson)
	var c cachorro
	fmt.Println(c)
	if erro := json.Unmarshal([]byte(cachorroEmJson), &c); erro != nil {
		log.Fatalln(erro)
	}
	fmt.Println(c)
	cachorro2EmJSON := `{"nome":"Draven", "raca":"Husky"}`
	var c2 map[string]string
	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil {
		log.Fatalln(erro)
	}
	fmt.Println(c2)
}
