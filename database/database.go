package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func testError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	urlConnection := "golang:golang@/golang?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", urlConnection)
	testError(erro)
	defer db.Close()

	erro = db.Ping()
	testError(erro)
	fmt.Println("Connected successfully")

	lines, erro := db.Query("select * from golang")
	testError(erro)
	defer lines.Close()
}
