package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello World")
	auxiliar.Escrever()
	auxiliar.Escrever2()
	erro := checkmail.ValidateFormat("humberto@gmail.com")
	fmt.Println(erro)
}
