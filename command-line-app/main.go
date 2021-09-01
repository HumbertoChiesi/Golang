package main

import (
	"command-line/app"
	"log"
	"os"
)

func main() {
	aplicacao := app.Gerar()

	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
