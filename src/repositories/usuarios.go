package repositories

import (
	"api/src/modelos"
	"database/sql"
)

type usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

func (repository usuarios) Criar(usuario modelos.Usuario) (uint64, error) {

	statement, erro := repository.db.Prepare("INSERT INTO usuarios (name, nick, email, password) VALUES (?,?,?,?);")

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	result, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)

	if erro != nil {
		return 0, erro
	}

	id, erro := result.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint64(id), nil
}
