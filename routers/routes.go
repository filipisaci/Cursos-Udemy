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

	api := App.Group("/v1")
	api.POST("/insert", c.Inserir)

}
