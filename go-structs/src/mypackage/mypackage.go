package mypackage

import "fmt"

type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

func PrintMessage(text string) {
	fmt.Println("Mensaje publico")
}

func printMessagel(text string) {
	fmt.Println("Mensaje publico")
}
