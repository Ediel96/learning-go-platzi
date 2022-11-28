package main

import "fmt"

func main() {

	helloMassage := "Hello"
	WorldMessage := "World"

	//Println(
	fmt.Println(helloMassage, WorldMessage)

	//Printf
	nombre := "Platzi"
	cursos := 500

	//%s tipos string
	//%d tipos entero
	//\n salto de linea
	fmt.Printf("%s tiene mas de %d cursos \n", nombre, cursos)

	//%v da igual que tipo de dato
	fmt.Printf("%v tiene mas de %v cursos \n", nombre, cursos)

	//Sprintf agregando mensage
	message := fmt.Sprintf("%s tiene mas de %d cursos \n", nombre, cursos)
	fmt.Printf(message)

	//tipo de dato
	fmt.Printf("helloMessage : %T \n", helloMassage)
	fmt.Printf("cursos : %T \n", cursos)

}
