package route

import (
	"usuarios/controller"

	"github.com/labstack/echo"
)

// App é um instância de echo
var App *echo.Echo

func init() {
	App = echo.New()

	App.GET("/", controller.Home)
	App.GET("/add", controller.Add)
	App.GET("/atualizar/:id", controller.Update)

	app := App.Group("/v1")
	app.POST("/insert", controller.Inserir)
	app.DELETE("/delete/:id", controller.Deletar)
	app.PUT("/update/:id", controller.Atualizar)
}
