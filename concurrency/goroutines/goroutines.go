package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Bom dia")
	escrever("Boa noite")
}

func escrever(txt string) {
	for {
		time.Sleep(time.Second * 2)
		fmt.Println(txt)
	}
}
