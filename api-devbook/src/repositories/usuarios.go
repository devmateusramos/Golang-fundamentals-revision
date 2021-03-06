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

func (repositorioUsuarios Usuarios) BuscarPorId(ID uint64) (models.Usuario, error) {
	linhas, err := repositorioUsuarios.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = ?", ID,
	)
	if err != nil {
		return models.Usuario{}, err
	}
	defer linhas.Close()

	var usuario models.Usuario

	if linhas.Next() {
		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return models.Usuario{}, err
		}
	}

	return usuario, nil
}

func (repositorioUsuarios Usuarios) Atualizar(ID uint64, usuario models.Usuario) error {
	statement, err := repositorioUsuarios.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); err != nil {
		return err
	}

	return nil
}

func (repositorioUsuarios Usuarios) Deletar(ID uint64) error {
	statement, err := repositorioUsuarios.db.Prepare(
		"delete from usuarios where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (repositorioUsuarios Usuarios) BuscarPorEmail(email string) (models.Usuario, error) {
	linha, err := repositorioUsuarios.db.Query("select id, senha from usuarios where email = ?", email)
	if err != nil {
		return models.Usuario{}, err
	}

	var usuario models.Usuario
	if linha.Next() {
		if err = linha.Scan(&usuario.ID, &usuario.Senha); err != nil {
			return models.Usuario{}, err
		}
	}

	return usuario, nil
}

// Seguir permite que um usu??rio siga o outro
func (repositorio Usuarios) Seguir(usuarioID, seguidorID uint64) error {
	statement, err := repositorio.db.Prepare(
		"insert ignore into seguidores (usuario_id, seguidor_id) values(?, ?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(usuarioID, seguidorID); err != nil {
		return err
	}

	return nil
}

func (repositorio Usuarios) PararDeSeguir(usuarioID, seguidorID uint64) error {
	statement, err := repositorio.db.Prepare("delete from seguidores where usuario_id = ? and seguidor_id = ?")
	if err != nil {
		return err
	}

	if _, err = statement.Exec(usuarioID, seguidorID); err != nil {
		return err
	}

	return nil
}

// BuscarSeguidores traz todos os seguidores de um usu??rio
func (repositorio Usuarios) BuscarSeguidores(usuarioID uint64) ([]models.Usuario, error) {
	linhas, err := repositorio.db.Query(
		"select u.id, u.nome, u.nick, u.email, u.criadoEm from usuarios u inner join seguidores s on u.id = s.seguidor_id where s.usuario_id = ?", usuarioID,
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
