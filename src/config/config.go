package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConexaoBanco = ""
	//	Porta onde a API vai estar rodando
	Porta = 0
	//	É a chave que vai ser usada para assinar o Token
	SecretKey []byte
)

func Carregar() {

	var erro error

	if erro = godotenv.Load(".env"); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))

	if erro != nil {
		Porta = 5000
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	fmt.Println("Porta configurada:", Porta)
	fmt.Println("Url DB:", StringConexaoBanco)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

}
