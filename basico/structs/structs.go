package main

import "fmt"

type usuario struct {
	idade int
	nome  string
	cpf   string
}

func main() {
	var p1 usuario
	p1.cpf = "40193841-13"
	p1.nome = "Humberto Chiesi"
	p1.idade = 21
	fmt.Println(p1)

	p2 := usuario{18, "Larissa Santos", "301941851-45"}
	fmt.Println(p2)

	p3 := usuario{idade: 51}
	fmt.Println(p3)
}
