package controller

import (
	"net/http"

	"github.com/filipisaci/Go-Web-Udemy/model"
	"github.com/labstack/echo"
)

/*
  index da aplicação
  Para exportar esta função (ou seja, para que ela fique acessível de outro arquivo)
  é preciso que o primeiro caracter do nome da função fique em maisusculo, no caso o H
*/
func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func Inserir(c echo.Context) error {
	nome := c.FormValue("nome")
	email := c.FormValue("email")

	var usuario model.Usuarios
	usuario.Nome = nome
	usuario.Email = email

	if nome != "" && email != "" {
		if _, err := model.UserModel.Insert(usuario); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"mensagem": "Não foi possível adicionar o registro no banco de dados!",
			})
		}
		return c.JSON(http.StatusCreated, map[string]string{
			"mensagem": "Usuário cadastrado com sucesso!",
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"mensagem": "Os campos não podem ser nulos",
	})
}
