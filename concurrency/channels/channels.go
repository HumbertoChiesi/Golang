package main

import (
	"fmt"
)

func main() {
	canal := make(chan string)
	escrever("Bom dia", canal)

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

}

func escrever(txt string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- txt
		fmt.Println(txt)
	}

	close(canal)
}
