package model

import (
	"github.com/filipisaci/Go-Web-Udemy/lib"
)

/*
  Abaixo temos a representação da tabela do banco de dados
  Note que entre crase está sendo passado o nome do campo da tabela no db
  E também como será apresentado o campo no json
*/

type Usuarios struct {
	ID    int    `db:"id" json:"id"`
	Nome  string `db:"nome" json:"nome"`
	Email string `db:"email" json:"email"`
}

// foi passado o nome da tabela do banco de dados chamada usuarios
var UserModel = lib.Connection.Collection("usuarios")
