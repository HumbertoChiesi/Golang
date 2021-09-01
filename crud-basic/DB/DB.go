package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	connectionString := "goteste:golang@/golang?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", connectionString)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
