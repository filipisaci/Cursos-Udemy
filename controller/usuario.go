package controller

import (
	"net/http"
	"strconv"

	"github.com/filipisaci/Go-Web-Udemy/model"
	"github.com/labstack/echo"
)

/*
  index da aplicação
  Para exportar esta função (ou seja, para que ela fique acessível de outro arquivo)
  é preciso que o primeiro caracter do nome da função fique em maisusculo, no caso o H
*/

func Get(c echo.Context) error {
	data := map[string]interface{}{
		"title":   "Título",
		"content": "Conteúdo",
	}
	return c.Render(http.StatusOK, "test.html", data)
}

func Home(c echo.Context) error {
	var usuarios []model.Usuarios

	if err := model.UserModel.Find().All(&usuarios); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"mensagem": "Erro ao recuperar valores do banco de dados!",
		})
	}

	data := map[string]interface{}{
		"titulo":   "Listagem de usuários",
		"usuarios": usuarios,
	}
	return c.Render(http.StatusOK, "index.html", data)
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
		return c.Redirect(http.StatusFound, "/")
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"mensagem": "Os campos não podem ser nulos",
	})
}

func Deletar(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// mesma coisa que um select * from tabela where id = ?
	resultado := model.UserModel.Find("id=?", id)

	if count, _ := resultado.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"mensagem": "Usuário não encontrado!",
		})
	}

	if err := resultado.Delete(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"mensagem": "Erro ao tentar deletar usuario da base!",
		})
	}

	return c.JSON(http.StatusAccepted, map[string]string{
		"mensagem": "Usuário apagado com sucesso!",
	})
}

func Add(c echo.Context) error {
	return c.Render(http.StatusOK, "add.html", nil)
}

func Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	nome := c.FormValue("nome")
	email := c.FormValue("email")

	var usuario model.Usuarios
	usuario.ID = id
	usuario.Nome = nome
	usuario.Email = email

	/*
		var usuario = model.Usuarios{
			ID:    id,
			Nome:  nome,
			Email: email,
		}
	*/

	resultado := model.UserModel.Find("id=?", id)

	if count, _ := resultado.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"mensagem": "Usuário não encontrado!",
		})
	}

	if err := resultado.Update(usuario); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"mensagem": "Erro ao atualizar o registro no banco de dados!",
		})
	}

	return c.JSON(http.StatusAccepted, usuario)
}

func Atualizar(c echo.Context) error {
	var id, _ = strconv.Atoi(c.Param("id"))

	var usuario model.Usuarios

	resultado := model.UserModel.Find("id=?", id)

	if count, _ := resultado.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"mensagem": "Usuário não encontrado no banco de dados!",
		})
	}

	if err := resultado.One(&usuario); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"mensagem": "Usuário não encontrado no banco de dados!",
		})
	}

	var data = map[string]interface{}{
		"usuario": usuario,
	}

	return c.Render(http.StatusOK, "update.html", data)
}
