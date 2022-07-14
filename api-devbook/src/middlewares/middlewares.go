package middlewares

import (
	"api-devbook/src/autenticacao"
	"api-devbook/src/respostas"
	"fmt"
	"net/http"
)

// Logger escreve informações da requisição no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n %s %s %s \n", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Autenticar verifica se o usuário fazendo a requisição está autenticado.
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := autenticacao.ValidarToken(r); err != nil {
			respostas.Err(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
