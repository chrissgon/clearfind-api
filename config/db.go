package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectDB - Conecta ao Banco e Retorna Conexao
func ConnectDB() *sql.DB {
	var (
		db  *sql.DB
		err error
	)

	// Cria Conexao
	db, err = sql.Open("mysql", "admin:70413093@tcp(127.0.0.1:3306)/clearfind")

	if err != nil {
		panic(err.Error())
	}

	// Retorna Conexao
	return db
}
