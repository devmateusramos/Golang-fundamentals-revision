package controllers

import (
	"api-devbook/src/database"
	"api-devbook/src/models"
	"api-devbook/src/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var usuario models.Usuario
	if err = json.Unmarshal(corpoRequest, &usuario); err != nil {
		log.Fatalln(err)
	}

	db, err := database.Conectar()
	if err != nil {
		log.Fatalln(err)
	}

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	repositorio.Criar(usuario)
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {

}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {

}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {

}
