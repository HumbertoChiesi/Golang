package main

import (
	"fmt"
)

func calculo(n1, n2 int) (soma, subtracao int) {
	soma = n1 + n2
	if n2 > n1 {
		n1, n2 = n2, n1
	}
	subtracao = n1 - n2
	return
}

func main() {
	soma, subtracao := calculo(1, 21)
	fmt.Println(soma, subtracao)
}
