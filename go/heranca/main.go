package main

import "fmt"

// Pessoa é um tipo
type Pessoa struct {
	nome   string
	idade  string
	altura float32
}

func (p *Pessoa) narra() string {
	return fmt.Sprintf("%s foi na padaria e comprou um pão...", p.nome)
}

func main() {
	murilo := Pessoa{
		nome:   "Murilo Locatelli",
		idade:  "34",
		altura: 1.81,
	}

	fmt.Println(murilo.narra())
}
