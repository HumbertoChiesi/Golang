package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2) //+2

	go func() {
		escrever("Bom dia")
		waitGroup.Done() //-1
	}()
	go func() {
		escrever("Boa noite")
		waitGroup.Done() //-1
	}()

	waitGroup.Wait() //waiting the waiting groups to end
}

func escrever(txt string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 2)
		fmt.Println(txt)
	}
}
