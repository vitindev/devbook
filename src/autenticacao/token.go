package autenticacao

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

/*
JWT - PEGA AS INFORÇÕES E VAI GERAR UMA STRING CHAMADA DE TOKEN, QUE VAI CONTER PERMISSÕES DO USER/AUTENTICAÇÃO NA API
*/
func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString(config.SecretKey)
}

// Validar Token verifica se o token passado na requisição é valido
func ValidarToken(r *http.Request) error {

	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)

	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token inválido")
}

func extrairToken(r *http.Request) string {

	token := r.Header.Get("Authorization")
	splitToken := strings.Split(token, " ")

	if len(splitToken) == 2 {
		return splitToken[1]
	}

	return ""
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
