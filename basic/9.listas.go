package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es un palindromo")
	}
}

func main() {
	var array [4]int

	array[0] = 1
	array[1] = 2

	fmt.Println(array, len(array))

	//len = cuantos elementos en un array

	//Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	//Metodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	//append
	slice = append(slice, 7)

	//Append nueva list
	newSlice := []int{8, 6}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	sliceString := []string{"Hola", "que", "haces"}

	for i, valor := range sliceString {
		fmt.Println(i, valor)
	}

	// ama
	// amor a roma
	var palabra string = "Ama"
	fmt.Println(&palabra)
	minus := strings.ToLower(palabra)
	isPalindromo(minus)
}
