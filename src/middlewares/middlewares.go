package middlewares

import (
	"api/src/autenticacao"
	"api/src/responses"
	"log"
	"net/http"
)

/*
Camada que vai ficar entre a requisição e a respostas
*/

func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if erro := autenticacao.ValidarToken(r); erro != nil {
			responses.Erro(w, http.StatusUnauthorized, erro)
			return
		}

		proximaFuncao(w, r)
	}
}
