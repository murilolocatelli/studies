package model

import (
	"usuarios/lib"
)

// Usuario representa a tabela de usuários na base de dados
type Usuario struct {
	ID    int    `db:"id" json:"id"`
	Nome  string `db:"nome" json:"nome"`
	Email string `db:"email" json:"email"`
}

// UsuarioModel recebe a tabela da base de dados
var UsuarioModel = lib.Sess.Collection("usuario")
