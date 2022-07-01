package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idadeMasPodeSerQualquerNome"`
}

func main() {
	//json.Marshal()   // Usado para converter map e struc pra json
	//json.Unmarshal() // Transformar de json pra map/struct
	cachorro1 := cachorro{"Draven", "Husky", 2}
	fmt.Println(cachorro1)
	cachorro1EmJson, erro := json.Marshal(cachorro1)
	if erro != nil {
		log.Fatalln(erro)
	}
	fmt.Println(cachorro1EmJson)
	fmt.Println(bytes.NewBuffer(cachorro1EmJson))
	cachorroEmMap := map[string]string{
		"nome":  "Lucian",
		"raça":  "Pastor Alemão",
		"idade": "3 sim em string",
	}
	fmt.Println(cachorroEmMap)
	cachorroMapToJSON, erro := json.Marshal(cachorroEmMap)
	if erro != nil {
		log.Fatalln(erro)
	}
	fmt.Println(cachorroMapToJSON)
	fmt.Println(bytes.NewBuffer(cachorroMapToJSON))
}
