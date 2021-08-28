package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int16 = 120
	fmt.Println(numero)

	var numero2 uint32 = 122
	fmt.Println(numero2)

	//alias/ INT32 = RUNE
	var numero3 rune = 1312
	fmt.Println(numero3)

	//alias/ Byte = int8
	var numero4 byte = 255
	fmt.Println(numero4)

	var numeroReal float32 = 3120.12
	numeroReal2 := 12.012391
	fmt.Println(numeroReal + float32(numeroReal2))

	char := 'A'
	fmt.Println(char)

	var booleano bool = true
	fmt.Println(booleano)

	var erro error = errors.New("Deu erro aí amigão")
	fmt.Println(erro)
}
