package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es uno")
	} else {
		fmt.Println("No es uno")
	}

	if valor1 == 0 && valor2 == 2 {
		fmt.Println("Es verdad OR")
	}

	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Es verdad OR")
	}

	//Convertir testo a numero
	value, err := strconv.Atoi("43")
	// Link de tipo de conversiones = https://pkg.go.dev/strconv

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Value:", value)

}
