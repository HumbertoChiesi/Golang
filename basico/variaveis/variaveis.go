package main

import (
	"fmt"
)

func main() {
	var variavel string = "variavel string"
	variavelImplicita := " variavel com declaracao implicita"
	fmt.Println(variavel + variavelImplicita)

	var (
		v1 string = "ifdajmsf"
		v2 string = "fknasfm"
	)
	fmt.Println(v1 + v2)

	v3, v4 := "v3 ", "v4 "
	fmt.Println(v3 + v4)

	const constante string = "variavel constante"
	fmt.Println(constante)

	v3, v4 = v4, v3
	fmt.Println(v3 + v4)
}
