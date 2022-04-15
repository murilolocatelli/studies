package main

import "fmt"

func soma(numero1 int, numero2 int) int {
	return numero1 + numero2
}

func main() {
	numero1 := 10
	numero2 := 20
	numero3 := 30
	numero4 := 40

	resultado := soma(numero1, numero2)
	resultado2 := soma(numero3, numero4)

	resultado *= 2

	fmt.Println(resultado)
	fmt.Println(resultado2)
}
