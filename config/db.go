package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// ConnectDB - Conecta ao Banco e Retorna Conexao
func ConnectDB() *sql.DB {
	var (
		db  *sql.DB
		err error
	)

	const (
		host     = "localhost"
		port     = 5432
		user     = "root"
		password = "12345"
		dbname   = "clearfind"
	)

	// Cria Conexao
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()

	if err != nil {
		panic(err.Error())
	}

	// Retorna Conexao
	return db
}
