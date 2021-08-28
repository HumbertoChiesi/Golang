package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	soma := somar(10, 45)
	fmt.Println(soma)

	var f = func(n int) (int, string) {
		if n%2 == 0 {
			return n % 2, "é divisível por 2"
		} else {
			return n % 2, "não é divisível por 2"
		}
	}

	n := 1127
	resto, txtDivisivel := f(n)
	fmt.Printf("%d %s, resto = %d", n, txtDivisivel, resto)
}
