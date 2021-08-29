package main

import "fmt"

func main() {
	canal := make(chan string, 2)
	canal <- "Bom dia"
	canal <- "Boa noite"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem, mensagem2)
}
