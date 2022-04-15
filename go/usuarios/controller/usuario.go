package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"usuarios/model"
)

// Home é a página inicial
func Home(c echo.Context) error {
	var usuarios []model.Usuario

	if err := model.UsuarioModel.Find().All(&usuarios); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"mensagem": "Erro ao tentar recuperar os dados!",
		})
	}

	data := map[string]interface{}{
		"titulo":   "Listagem de usuários",
		"usuarios": usuarios,
	}

	return c.Render(http.StatusOK, "index.html", data)
}

// Add é a página para criação de usuário
func Add(c echo.Context) error {
	return c.Render(http.StatusOK, "add.html", nil)
}

// Update é a página para atualização de usuário
func Update(c echo.Context) error {
	usuarioID, _ := strconv.Atoi(c.Param("id"))

	var usuario model.Usuario

	resultado := model.UsuarioModel.Find("id=?", usuarioID)

	if count, _ := resultado.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"mensagem": "Não foi possível encontrar o usuário!",
		})
	}

	if err := resultado.One(&usuario); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"mensagem": "Não foi possível encontrar o usuário!",
		})
	}

	var data = map[string]interface{}{
		"usuario": usuario,
	}

	return c.Render(http.StatusOK, "update.html", data)
}

// Inserir é a função resposável por salvar o usuário na base de dados
func Inserir(c echo.Context) error {
	nome := c.FormValue("nome")
	email := c.FormValue("email")

	var usuario model.Usuario
	usuario.Nome = nome
	usuario.Email = email

	if nome != "" && email != "" {
		if _, err := model.UsuarioModel.Insert(usuario); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"mensagem": "Não foi possível adicionar o registro no banco!",
			})
		}

		/*return c.JSON(http.StatusCreated, map[string]string{
			"mensagem": "O registro foi salvo com sucesso!",
		})*/

		return c.Redirect(http.StatusFound, "/")
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"mensagem": "Os campos precisam ser informados!",
	})
}

// Deletar é a função resposável por deletar o usuário na base de dados
func Deletar(c echo.Context) error {
	usuarioID, _ := strconv.Atoi(c.Param("id"))

	resultado := model.UsuarioModel.Find("id=?", usuarioID)

	if count, _ := resultado.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"mensagem": "Não foi possível encontrar o usuário!",
		})
	}

	if err := resultado.Delete(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"mensagem": "Não foi possível deletar o usuário!",
		})
	}

	return c.JSON(http.StatusAccepted, map[string]string{
		"mensagem": "Usuário deletado com sucesso!",
	})
}

// Atualizar é a função resposável por atualizar o usuário na base de dados
func Atualizar(c echo.Context) error {
	usuarioID, _ := strconv.Atoi(c.Param("id"))
	nome := c.FormValue("nome")
	email := c.FormValue("email")

	var usuario = model.Usuario{
		ID:    usuarioID,
		Nome:  nome,
		Email: email,
	}

	resultado := model.UsuarioModel.Find("id=?", usuarioID)

	if count, _ := resultado.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"mensagem": "O usuário não existe!",
		})
	}

	if err := resultado.Update(usuario); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"mensagem": "Erro ao tentar atualizar o registro!",
		})
	}

	return c.JSON(http.StatusAccepted, usuario)
}
