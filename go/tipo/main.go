package main

import "fmt"

// Pessoa é um tipo
type Pessoa string

func main() {
	var nome string
	var idade int
	var dinheiro float32
	var verificar bool
	var murilo Pessoa

	nome = "Murilo Locatelli"
	idade = 34
	dinheiro = 50.22
	verificar = true
	murilo = "Locatelli"

	fmt.Println(nome)
	fmt.Println(idade)
	fmt.Println(dinheiro)
	fmt.Println(verificar)
	fmt.Println(murilo)
}
