package main

import "fmt"

func soma(numeros ...int) int {
	somatoria := 0
	for _, n := range numeros {
		somatoria += n
	}
	return somatoria
}

func main() {
	fmt.Println(soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
