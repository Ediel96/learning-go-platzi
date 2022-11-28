package main

import "fmt"

func main() {

	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	//si condicion

	value := 50
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es mayor a 0")
	default:
		fmt.Println("No codicion")
	}
}
