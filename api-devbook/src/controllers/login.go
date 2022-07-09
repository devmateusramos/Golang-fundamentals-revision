package controllers

import (
	"api-devbook/src/autenticacao"
	"api-devbook/src/database"
	"api-devbook/src/models"
	"api-devbook/src/repositories"
	"api-devbook/src/respostas"
	"api-devbook/src/seguranca"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Err(w, http.StatusUnprocessableEntity, err)
	}

	var usuario models.Usuario
	if err := json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		respostas.Err(w, http.StatusBadRequest, err)
	}

	db, err := database.Conectar()
	if err != nil {
		respostas.Err(w, http.StatusInternalServerError, err)
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	usuarioSalvoNoBanco, err := repositorio.BuscarPorEmail(usuario.Email)
	if err != nil {
		respostas.Err(w, http.StatusInternalServerError, err)
		return
	}

	if err = seguranca.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha); err != nil {
		respostas.Err(w, http.StatusUnauthorized, err)
		return
	}

	token, err := autenticacao.CriarToken(usuarioSalvoNoBanco.ID)
	if err != nil {
		respostas.Err(w, http.StatusInternalServerError, err)
		return
	}
	w.Write([]byte(token))
}
