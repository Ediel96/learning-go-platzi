package main

import (
	"fmt"
	pk "learn-go/src/mypackage"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrrari"
	fmt.Println(myCar)
	pk.PrintMessage("Hello Platzi")

	// No se puede
	// pk.printMessagel("")
}
