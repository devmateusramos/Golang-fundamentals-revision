package repositories

import (
	"api-devbook/src/models"
	"database/sql"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (usuarios Usuarios) Criar(usuario models.Usuario) (uint64, error) {
	return 0, nil
}
