package main

import (
	pongor "github.com/MarcusMann/pongor"
	r "github.com/filipisaci/Go-Web-Udemy/routers"
	"github.com/labstack/echo/middleware"
)

func main() {
	// e é uma instancia da lib echo, utilizada para trabalhar com metodos http
	e := r.App
	// p é uma instancia de pongo2 que é utilizada para criar templates html
	p := pongor.GetRenderer(pongor.PongorOption{Directory: "view", Reload: true})

	e.Renderer = p

	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":3000"))
}
