package main

import "fmt"

func show(nome string) {
	nome = "Locatelli"
}

func show2(nome *string) {
	*nome = "Locatelli"
}

func main() {
	var nome = "Murilo"

	ref := &nome

	fmt.Println(nome)
	fmt.Println(ref)
	fmt.Println(*ref)

	show(nome)

	fmt.Println(nome)

	show2(&nome)

	fmt.Println(nome)
}
