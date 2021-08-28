package main

import "fmt"

func f1() {
	fmt.Println("funcao 1")
}

func f2() {
	fmt.Println("funcao 2")
}

func main() {
	defer f1()
	f2()

	fmt.Println("defer chama a funcao f1 apenas quando a funcao main for se encerrar")
}
