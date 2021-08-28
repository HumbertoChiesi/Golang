package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
	sexo  string
}

func main() {
	nomes := []string{"Humberto", "Vinicius", "Lais", "Humberto"}

	var p pessoa
	p.idade = 21
	p.nome = "Humberto"
	p.sexo = "Masculino"

	for indice, valor := range nomes {
		fmt.Println(indice, valor)
	}
}
