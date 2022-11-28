package main

import "fmt"

func main() {

	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println("pi : ", pi)
	fmt.Println("pi2 : ", pi2)

	//Declaracion de variables
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("are cuadrado : ", areaCuadrado)

	x := 10
	y := 50

	result := x + y

	fmt.Println("suma", result)

	result = x - y

	fmt.Println("resta ", result)

	result = x * y

	fmt.Println("resta ", result)

	result = x / y

	fmt.Println("divicion ", result)

	result = x % y

	fmt.Println("residuo ", result)

	//Incremental
	x++
	fmt.Println("incremental : ", x)

	// Decremental
	y++
	fmt.Println("dremental : ", y)

}
