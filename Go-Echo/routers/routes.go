package routers

import (
	c "github.com/filipisaci/Go-Web-Udemy/controller"
	"github.com/labstack/echo"
)

var App *echo.Echo

//funcao construtora
func init() {
	App = echo.New()

	App.GET("/", c.Home)
	App.GET("/add", c.Add)
	App.GET("/atualizar/:id", c.Atualizar)

	api := App.Group("/v1")
	api.POST("/insert", c.Inserir)
	api.GET("/user", c.Get)
	api.DELETE("/delete/:id", c.Deletar)
	api.PUT("/update/:id", c.Update)
}
