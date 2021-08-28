package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fezAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Humberto", 17}
	fmt.Println(usuario1)

	if usuario1.maiorDeIdade() {
		fmt.Printf("Usuario %s é maior de idade\n", usuario1.nome)
	} else {
		fmt.Printf("Usuario %s é menor de idade\n", usuario1.nome)
	}

	usuario1.fezAniversario()
	fmt.Println(usuario1.idade)

	if usuario1.maiorDeIdade() {
		fmt.Printf("Usuario %s é maior de idade\n", usuario1.nome)
	} else {
		fmt.Printf("Usuario %s é menor de idade\n", usuario1.nome)
	}
}
