package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Bom dia", canal)

	for {
		mensagem, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}

}

func escrever(txt string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- txt
		time.Sleep(time.Second)
	}

	close(canal)
}
