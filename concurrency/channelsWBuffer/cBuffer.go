package main

import "fmt"

func main() {
	canal := make(chan string, 1)
	canal <- "Bom dia"
	canal <- "Boa noite"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem, mensagem2)
}
