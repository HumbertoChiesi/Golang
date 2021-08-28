package main

import "fmt"

func main() {
	var numero1 int = 120
	var numero2 int = numero1
	var ponteiroNumero *int = &numero1

	fmt.Println(numero1, numero2, *ponteiroNumero)

	numero1 += 300

	fmt.Println(numero1, numero2, *ponteiroNumero)
}
