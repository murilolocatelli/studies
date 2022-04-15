package main

import "fmt"

func soma(numeros ...int) int {
	result := 0

	for _, numero := range numeros {
		result += numero
	}

	return result
}

func main() {
	resultado := soma(1, 2, 3)

	fmt.Println(resultado)
}
