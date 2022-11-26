package main

import "fmt"

func funcionNormal(message string) {
	fmt.Println(message)
}

func tripeArgument(a, b int, c string) {
	fmt.Println(a, b, c)

}

func retunValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (b, c int) {
	return a, a * 2
}

func main() {
	funcionNormal("Hola mundo")
	tripeArgument(1, 49, "string")

	value := retunValue(2)
	fmt.Println("value : ", value)

	value1, value2 := doubleReturn(2)

	value3, _ := doubleReturn(2) // solo solo incluir un solo valor
	fmt.Println("double return : ", value1, value2, value3)
}
