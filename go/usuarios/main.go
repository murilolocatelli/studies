package main

import (
	r "usuarios/route"

	"github.com/labstack/echo/middleware"

	"github.com/MarcusMann/pongor"
)

func main() {
	r.App.Use(middleware.Logger())

	p := pongor.GetRenderer(pongor.PongorOption{Reload: true})
	p.Directory = "view"

	r.App.Renderer = p

	r.App.Start(":3000")
}
