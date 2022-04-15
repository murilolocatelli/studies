package main

import "fmt"

func main() {
	mapa := map[string]string{
		"nome":  "Murilo Locatelli",
		"idade": "34",
	}

	mapa2 := make(map[string]string)

	mapa2["nome"] = "Dolores"
	mapa2["idade"] = "50"

	fmt.Println(mapa["nome"])
	fmt.Println(mapa2["idade"])
}
