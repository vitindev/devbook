package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

/*
Basicamente no URI temos os parametros que são definidos via localhost:5000/usuarios/{parametro},
e também temos as query que são definidas via localhost:5000/usuarios?usuario=Fernando
*/

/*

Gerando uma SECRET KEY simples:

func init() {

	chave := make([]byte, 64)

	if _, erro := rand.Read(chave); erro != nil {
		log.Fatal(erro)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(chave)
	fmt.Println(stringBase64)

}
*/

func main() {

	config.Carregar()

	fmt.Println("Rodando API!")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
