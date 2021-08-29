package main

import "fmt"

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for n := 0; n < 45; n++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

func worker(t <-chan int, r chan<- int) {
	for numeros := range t {
		r <- fibonacci(numeros)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
