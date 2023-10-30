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

func main() {

	config.Carregar()

	fmt.Println("Rodando API!")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
