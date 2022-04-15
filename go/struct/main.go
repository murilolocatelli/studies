package main

import "fmt"

type modelo string

// Pessoa é um tipo
type Pessoa struct {
	nome   modelo
	idade  string
	altura float32
}

func show(p Pessoa) {
	fmt.Println(p.altura)
}

func main() {
	murilo := Pessoa{}
	murilo.nome = "Murilo Locatelli"
	murilo.idade = "34"
	murilo.altura = 1.81

	show(murilo)
}
