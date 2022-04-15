package lib

import (
	"log"

	"upper.io/db.v3"

	"upper.io/db.v3/mysql"
)

var configuracao = mysql.ConnectionURL{
	Host:     "localhost",
	User:     "root",
	Password: "teste123456",
	Database: "membros",
}

// Sess variável que faz a conexão com a base de dados
var Sess db.Database

func init() {
	var err error

	Sess, err = mysql.Open(configuracao)

	if err != nil {
		log.Fatal(err.Error())
	}
}
