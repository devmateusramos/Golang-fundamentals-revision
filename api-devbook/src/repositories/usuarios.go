package repositories

import (
	"api-devbook/src/models"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorioUsuarios Usuarios) Criar(usuario models.Usuario) (uint64, error) {
	statement, err := repositorioUsuarios.db.
		Prepare("insert into usuarios (nome, nick, email, senha) values(?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return uint64(ultimoIDInserido), nil
}

func (repositorioUsuarios Usuarios) Buscar(nomeOuNick string) ([]models.Usuario, error) {
	nomeOuNickFormated := fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, err := repositorioUsuarios.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome like ? or nick like ?",
		nomeOuNickFormated, nomeOuNickFormated,
	)
	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario

		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}
