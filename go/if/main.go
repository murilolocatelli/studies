package main

import "fmt"

func main() {
	var code bool
	code = true

	if code == true {
		fmt.Println("Passou no teste")
	} else {
		fmt.Println("Não passou no teste")
	}

	if code {
		fmt.Println("Passou no teste")
	} else {
		fmt.Println("Não passou no teste")
	}

	if code2 := false; code2 {
		fmt.Println("Passou no teste")
	} else {
		fmt.Println("Não passou no teste")
	}
}
