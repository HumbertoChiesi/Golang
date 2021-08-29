package main

import (
	"fmt"
	"time"
)

func main() {
	c1, c2 := make(chan string), make(chan string)
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			c1 <- "Canal1"
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second * 2)
			c2 <- "Canal2"
		}
	}()
	for {

		select {
		case cMensagem1 := <-c1:
			fmt.Println(cMensagem1)
		case cMensagem2 := <-c2:
			fmt.Println(cMensagem2)
		}
	}
}
