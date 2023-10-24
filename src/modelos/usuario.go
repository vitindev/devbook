package modelos

import (
	"errors"
	"strings"
	"time"
)

//`json:"id,omitempty"` omitempty não passa quando tiver valor 0

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

func (usuario *Usuario) Preparar() error {

	if erro := usuario.validar(); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar() error {

	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	} else if usuario.Nick == "" {
		return errors.New("O nick é obrigatório e não pode estar em branco")
	} else if usuario.Email == "" {
		return errors.New("O email é obrigatório e não pode estar em branco")
	} else if usuario.Senha == "" {
		return errors.New("A senha é obrigatória e não pode estar em branco")
	}

	return nil

}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
