package main

import "fmt"

type pessoa struct {
	nome   string
	idade  int
	altura float32
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"Humberto Chiesi", 21, 1.80}
	fmt.Println(p1)

	e := estudante{p1, "Ciêcina da Computação", "PUC-SP"}
	fmt.Println(e)
}
