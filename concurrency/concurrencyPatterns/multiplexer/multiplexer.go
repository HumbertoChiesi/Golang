package main

import (
	"fmt"
	"time"
)

func main() {
	canal := multiplexar(escrever("Bom dia"), escrever("Boa noite"))

	for i := 0; i < 15; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexar(c1, c2 <-chan string) <-chan string {
	canalSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-c1:
				canalSaida <- mensagem
			case mensagem := <-c2:
				canalSaida <- mensagem
			}
		}
	}()

	return canalSaida
}

func escrever(txt string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s ", txt)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
