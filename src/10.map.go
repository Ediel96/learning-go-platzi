package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	//Recorrer map
	for i, v := range m {
		fmt.Println(i, v)
	}

	//Encontrar valir
	value, ok := m["Jose"]
	fmt.Println(value, ok)
}
