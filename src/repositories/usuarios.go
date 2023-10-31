package repositories

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
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

func (repository usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {

	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, erro := repository.db.Query("SELECT * FROM usuarios WHERE name LIKE(?) OR nick LIKE(?);", nomeOuNick, nomeOuNick)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {

		var usuario modelos.Usuario

		if erro = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.Senha, &usuario.CriadoEm); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repository usuarios) BuscarPorID(usuarioId uint64) (modelos.Usuario, error) {

	linha, erro := repository.db.Query("SELECT * FROM usuarios WHERE id=? LIMIT 1;", usuarioId)

	if erro != nil {
		return modelos.Usuario{}, erro
	}

	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {

		if erro = linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.Senha, &usuario.CriadoEm); erro != nil {
			return modelos.Usuario{}, erro
		}

	}

	return usuario, nil
}

func (repository usuarios) Atualizar(usuarioId uint64, usuario modelos.Usuario) error {

	statement, erro := repository.db.Prepare("UPDATE usuarios SET name=?, nick=?, email=? WHERE id=? LIMIT 1;")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuarioId); erro != nil {
		return erro
	}

	return nil
}

func (repository usuarios) Deletar(usuarioId uint64) error {

	statement, erro := repository.db.Prepare("DELETE FROM usuarios WHERE id=? LIMIT 1;")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(usuarioId); erro != nil {
		return erro
	}

	return nil
}
